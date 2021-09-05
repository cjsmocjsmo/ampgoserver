///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"fmt"
	"log"
	// "bytes"
	"time"
	
	"strconv"
	"math/rand"
	
	// "path"
	// "strings"
	// "io/ioutil"
	// "sort"
	"net/http"
	// "html/template"
	"encoding/hex"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/bson"
	"context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/globalsign/mgo"
	// "github.com/globalsign/mgo/bson"
	
	"github.com/cjsmocjsmo/ampgosetup"
)

type plist struct {
	PLName string              `bson:"PLName"`
	PLId   string              `bson:"PLId"`
	Songs  []map[string]string `bson:"Songs"`
}

type iMgfa struct {
	Album   string              `bson:"album"`
	PicPath string              `bson:"picPath"`
	Songs   []map[string]string `bson:"songs"`
}

type rAlbinfo struct {
	Songs   []map[string]string `bson:"songs"`
	HSImage string              `bson:"hsimage"`
}

type voodoo struct {
	Playlists []map[string]string `bson:"playlists"`
}

//ArtVIEW exported
type ArtVIEW struct {
	Artist   string              `bson:"artist"`
	ArtistID string              `bson:"artistID"`
	Albums   []map[string]string `bson:"albums"`
	Page     string              `bson:"page"`
	Idx      string              `bson:"idx"`
}

//Albview exported
type AlbvieW struct {
	Artist   string              `bson:"artist"`
	ArtistID string              `bson:"artistID"`
	Album    string              `bson:"album"`
	AlbumID  string              `bson:"albumID"`
	Songs    []map[string]string `bson:"songs"`
	Page     string              `bson:"page"`
	NumSongs string              `bson:"numsongs"`
	PicPath  string              `bson:"picPath"`
	Idx      string              `bson:"idx"`
}

type AmpgoRandomPlaylistData struct {
	PlayListName string `bson:"playlistname"`
	PlayListID string `bson:"playlistID"`
	PlayList []map[string]string `bson:"playlist"`
}

var OFFSET string = os.Getenv("AMPGO_OFFSET")

func RemoveLogFile(logtxtfile string) bool {
	// var logtxtfile string = os.Getenv("AMPGO_SERVER_LOG_PATH")
	var result bool
	_, err := os.Stat(logtxtfile)
    if err == nil {
        log.Printf("logfile %s exists removing", logtxtfile)
		os.Remove(logtxtfile)
		result = true
    } else if os.IsNotExist(err) {
        log.Printf("logfile %s does not exists", logtxtfile)
		result = true
    } else {
        log.Printf("logfile %s stat error: %v", logtxtfile, err)
		result = false
    }
	return result
}

func StartServerLogging() string {
	var logtxtfile string = os.Getenv("AMPGO_SERVER_LOG_PATH")
	result := RemoveLogFile(logtxtfile)
	if result {
		file, err := os.OpenFile(logtxtfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	} else {
		fmt.Println("Unable to setup logging")
	}
	return "Logging started"
}

func ServerCheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Println(msg)
		panic(err)
	}
}

// func sfdbCon() *mgo.Session {
// 	s, err := mgo.Dial("mongodb://db:27017/ampgodb")
// 	if err != nil {
// 		log.Println("Session creation dial error")
// 		log.Println(err)
// 	}
// 	log.Println("Session Connection to db established")
// 	return s
// }

func setUpHandler(w http.ResponseWriter, r *http.Request) {
	ampgosetup.Setup()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Setup Complete")
	//this needs work
	log.Println("Setup Complete")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello From Ampgo Home \n It works")
	log.Println("homeHandler is complete")
}


///////////////////////////////////////////////////////////////////////////////
//Artist Stuff
///////////////////////////////////////////////////////////////////////////////

func initArtistInfoHandler(w http.ResponseWriter, r *http.Request) {
	// limit, err := strconv.ParseInt(OFFSET, 10, 64)
	// ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{}}
	opts := options.Find()
	// opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0, "artist": 1, "artistID": 1, "albcount": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("artistview").Collection("artistview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initArtistInfo find has failed")
	var allartist []map[string]string
	if err = cur.All(context.TODO(), &allartist); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s this is allartist-", allartist)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allartist)
	log.Println("Init Artist Info Complete")

}


func ArtViewFindOne(db string, coll string, filtertype string, filterstring string) ArtVIEW {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results ArtVIEW
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil { log.Fatal(err) }
	return results
}




func albumsForArtist2Handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println("Starting albumsForArtistHandler")
	var artistid string = r.URL.Query().Get("selected")
	log.Printf("%s this is artistid", artistid)
	log.Printf("%T this is artistid type", artistid)
	allalbums := ArtViewFindOne("artistview", "artistview", "artistID", artistid)
	// filter := bson.D{{"artistID", artistid}}
	// opts := options.Find()
	// opts.SetProjection(bson.M{"_id": 0, "songs": 0})
	// limit, err := strconv.ParseInt(OFFSET, 10, 64)
	// ServerCheckError(err, "convert to int64 has failed")
	// filter := bson.D{{}}
	// opts := options.Find()
	// opts.SetLimit(int64(limit))
	// opts.SetProjection(bson.M{"_id": 0, "albums": 1})
	// client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	// defer ampgosetup.Close(client, ctx, cancel)
	// ServerCheckError(err, "MongoDB connection has failed")
	// coll := client.Database("artistview").Collection("artistview")
	// cur, err := coll.Find(context.TODO(), filter, opts)
	// ServerCheckError(err, "initArtistInfo find has failed")
	// var allartist ArtVIEW
	// if err = cur.All(context.TODO(), &allartist); err != nil {
	// 	log.Fatal(err)
	// }
	for _, a := range allalbums.Albums {
		log.Println(a)
	}
	log.Println(&allalbums.Albums)
	log.Printf("%T this is allartist type", &allalbums.Albums)
	log.Printf("%s this is allartist-", &allalbums.Albums)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbums.Albums)
	log.Println("Init Artist Info Complete")
}



// func albumsForArtistHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Starting albumsForArtistHandler")
// 	var artistid string = r.URL.Query().Get("selected")
// 	log.Printf("%s this is artistid", artistid)
// 	log.Printf("%T this is artistid type", artistid)
// 	filter := bson.D{{"artistID", artistid}}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0, "songs": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("albumview").Collection("albumview")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "initArtistInfo find has failed")
// 	var allalbum []map[string]string
// 	if err = cur.All(context.TODO(), &allalbum); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is allalbum-", allalbum)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allalbum)
// 	log.Println("Init Album Info Complete")
// }

func songsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting songsForAlbumHandler")
	var albumid string = r.URL.Query().Get("selected")
	log.Printf("%s this is albumid", albumid)
	log.Printf("%T this is albumid type", albumid)
	filter := bson.D{{"albumID", albumid}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "songsForAlbumHandler find has failed")
	var allalbum []map[string]string
	if err = cur.All(context.TODO(), &allalbum); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s this is allalbum-", allalbum)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbum)
	log.Println("songsForAlbumHandler Complete")
}

///////////////////////////////////////////////////////////////////////////////
//Album Stuff
///////////////////////////////////////////////////////////////////////////////



func initalbumInfoHandler(w http.ResponseWriter, r *http.Request) {
	// limit, err := strconv.ParseInt(OFFSET, 10, 64)
	// ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{}}
	opts := options.Find()
	// opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0, "Album": 1, "AlbumID": 1, "PicPath": 1}) //must be uppercase did not use a struct
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("tempdb2").Collection("artidpic")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initAlbumInfo find has failed")
	var allalbums []map[string]string
	if err = cur.All(context.TODO(), &allalbums); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s this is allalbums", allalbums)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbums)
	log.Println("Init albumsInfo Complete")
}


func initialsongInfoHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "ArtPipeline find has failed")
	var tv []map[string]string
	if err = cur.All(context.TODO(), &tv); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s this is tv", tv)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tv)
	log.Println("Initial Song Info Complete")
}






// func artistPageHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.D{}
// 	opts := options.Distinct()
// 	opts.SetMaxTime(2 * time.Second)
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("artistview").Collection("artistview")
// 	DD1, err2 := collection.Distinct(context.TODO(), "page", filter, opts)
// 	ServerCheckError(err2, "MongoDB distinct album has failed")
// 	var ARDist []string
// 	for _, DD := range DD1 {
// 		zoo := fmt.Sprintf("%s", DD)
// 		ARDist = append(ARDist, zoo)
// 	}
// 	sort.Strings(ARDist)
// 	log.Println("ArtistAlpha is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(ARDist)
// }

// func albumPageHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.D{}
// 	opts := options.Distinct()
// 	opts.SetMaxTime(2 * time.Second)
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("albumview").Collection("albumview")
// 	DD1, err2 := collection.Distinct(context.TODO(), "albumpage", filter, opts)
// 	ServerCheckError(err2, "MongoDB distinct album has failed")
// 	var ALDist []string
// 	for _, DD := range DD1 {
// 		zoo := fmt.Sprintf("%s", DD)
// 		ALDist = append(ALDist, zoo)
// 	}
// 	sort.Strings(ALDist)
// 	log.Println("AlbumAlpha is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(ALDist)
// }

// func titlePageHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.D{}
// 	opts := options.Distinct()
// 	opts.SetMaxTime(2 * time.Second)
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("maindb").Collection("maindb")
// 	DD1, err2 := collection.Distinct(context.TODO(), "titlepage", filter, opts)
// 	ServerCheckError(err2, "MongoDB distinct album has failed")
// 	var TDist []string
// 	for _, DD := range DD1 {
// 		zoo := fmt.Sprintf("%s", DD)
// 		TDist = append(TDist, zoo)
// 	}
// 	sort.Strings(TDist)
// 	log.Println("TitleAlpha is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(TDist)
// }

// func songInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	pagenum := r.URL.Query().Get("selected")
// 	limit, err := strconv.ParseInt(OFFSET, 10, 64)
// 	ServerCheckError(err, "convert to int64 has failed")
// 	filter := bson.D{{"titlepage", pagenum}}
// 	opts := options.Find()
// 	opts.SetLimit(int64(limit))
// 	opts.SetProjection(bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("maindb")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "ArtPipeline find has failed")
// 	var SIS []map[string]string
// 	if err = cur.All(context.TODO(), &SIS); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("SongInfo is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&SIS)
// }

// func albumInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	pagenum := r.URL.Query().Get("selected")
// 	limit, err := strconv.ParseInt(OFFSET, 10, 64)
// 	ServerCheckError(err, "convert to int64 has failed")
// 	filter := bson.D{{"albumpage", pagenum}}
// 	opts := options.Find()
// 	opts.SetLimit(int64(limit))
// 	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
// 	opts.SetProjection(b2)
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("albumview").Collection("albumview")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "AlbPipeline find has failed")
// 	var AI []AlbvieW
// 	if err = cur.All(context.TODO(), &AI); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("AlbumInfo is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&AI)
// }

// func artistInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	pagenum := r.URL.Query().Get("selected")
// 	limit, err := strconv.ParseInt(OFFSET, 10, 64)
// 	ServerCheckError(err, "convert to int64 has failed")
// 	filter := bson.D{{"page", pagenum}}
// 	opts := options.Find()
// 	opts.SetLimit(int64(limit))
// 	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
// 	opts.SetProjection(b2)
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("artistview").Collection("artistview")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "ArtPipeline find has failed")
// 	var ARTI []ArtVIEW
// 	if err = cur.All(context.TODO(), &ARTI); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("ArtistInfo is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&ARTI)
// }

// func imageSongsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Starting imageSongsForAlbumHandler")
// 	albumid := r.URL.Query().Get("selected")
// 	log.Printf("%s this is albumid", albumid)

// 	limit, err := strconv.ParseInt(OFFSET, 10, 64)
// 	ServerCheckError(err, "ParseInt has failed")
// 	filter := bson.D{{"albumID", albumid}}
// 	opts := options.Find()
// 	opts.SetLimit(int64(limit))
// 	b2 := bson.M{"_id": 0, "album": 1, "picPath": 1, "songs": 1}
// 	opts.SetProjection(b2)
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("albumview").Collection("albumview")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "imageSongsForAlbumHandler has failed")
// 	var sfora []iMgfa
// 	if err = cur.All(context.TODO(), &sfora); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("this is sfora")
// 	log.Println(sfora)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&sfora)
// }

func randomPicsHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "index": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("coverart").Collection("coverart")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "imageSongsForAlbumHandler has failed")
	var indexliststring []map[string]string
	if err = cur.All(context.TODO(), &indexliststring); err != nil {
		log.Fatal(err)
	}
	var num_list []int
	for _, idx := range indexliststring {
		indexx := idx["index"]
		index1, _ := strconv.Atoi(indexx)
		num_list = append(num_list, index1)
	}
	Shuffle(num_list)
	log.Println(num_list)
	var randpics []string
	for _, f := range num_list[:12] {
		log.Printf("f type: %T", f)
		log.Printf("f: %s", f)
		ff := strconv.Itoa(f)
		log.Printf("ff type %T", ff)
		log.Printf("ff type %s", ff)
		filter := bson.D{{"index", ff}}
		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
		defer ampgosetup.Close(client, ctx, cancel)
		ampgosetup.CheckError(err, "MongoDB connection has failed")
		collection := client.Database("coverart").Collection("coverart")
		var rpics map[string]string
		err = collection.FindOne(context.Background(), filter).Decode(&rpics)
		if err != nil { log.Fatal(err) }
		randpics = append(randpics, rpics["imagehttpaddr"])
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randpics)
}

// func statsHandler(w http.ResponseWriter, r *http.Request) {
// 	// ST := ampgolib.GStats()
// 	ses := sfdbCon()
// 	defer ses.Close()
// 	STATc := ses.DB("goampgo").C("dbstats")
// 	b1 := bson.M{"_id": 0}
// 	var st map[string]string
// 	err := STATc.Find(nil).Select(b1).One(&st)
// 	if err != nil {
// 		log.Println("stats has fucked up")
// 		log.Println(err)
// 	}
// 	log.Println("GStats is complete")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&st)
// }



// func ramdomAlbumPicPlaySongHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("sid")
// 	rapp := ampgolib.RamdomAlbPicPlay(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(rapp)
// }

// func pathArtHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("selected")
// 	pa := ampgolib.PathArt(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(pa)
// }



// func songSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("searchval")
// 	artS := ampgolib.SongSearch(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(artS)
// }

// func albumSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("albsearchval")
// 	albS := ampgolib.AlbumSearch(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(albS)
// }

// func artistSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("artsearchval")
// 	artS := ampgolib.ArtistSearch(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(artS)
// }

// func addPlayListNameToDBHandler(w http.ResponseWriter, r *http.Request) {
// 	qu := r.URL.Query().Get("playlistname")
// 	aplntdb := ampgolib.AddPlayListNameToDB(qu)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(aplntdb)
// }

// func addSongsToPlistDBHandler(w http.ResponseWriter, r *http.Request) {
// 	sn := r.URL.Query().Get("songname")
// 	sid := r.URL.Query().Get("songid")
// 	plid := r.URL.Query().Get("playlistid")
// 	ampgolib.AddSongsToPlistDB(sn, sid, plid)
// }

// func allPlaylistSongsFromDBHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	soho := ampgolib.AllPlaylistSongsFromDB(plid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(soho)
// }

// func createPlayerPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	apsf := ampgolib.CreatePlayerPlaylist(plid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(apsf)
// }


// func deletePlaylistFromDBHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	dpl := ampgolib.DeletePlaylistFromDB(plid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(dpl)
// }

// func deleteSongFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plname := r.URL.Query().Get("playlistname")
// 	songid := r.URL.Query().Get("delsongid")
// 	dsfp := ampgolib.DeleteSongFromPlaylist(plname, songid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(dsfp)
// }









///////////////////////////////////////////////////////////////////////////////
//Playlist Stuff
///////////////////////////////////////////////////////////////////////////////

func addPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	plname := r.URL.Query().Get("name")
	plID, _ := UUID()
	var emptymap []map[string]string
	var plzz AmpgoRandomPlaylistData
	plzz.PlayListName = plname
	plzz.PlayListID = plID
	plzz.PlayList = emptymap
	log.Println("This is plzz")
	log.Println(plzz)
	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
	ServerCheckError(err3, "Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", &plzz)
	ServerCheckError(err2, "plz insertion has failed")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&plzz)
}

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
	   randIndex := r.Intn(n)
	   slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

func createRandomPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	plname := r.URL.Query().Get("name")
	plc := r.URL.Query().Get("songcount")
	log.Printf("planame: %s", plname)
	log.Printf("plc: %s", plc)
	plcount, _ := strconv.Atoi(plc)
	plID, _ := UUID()
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "idx": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "allIdx has failed")
	var indexlist []map[string]string
	if err = cur.All(context.TODO(), &indexlist); err != nil {
		log.Fatal(err)
	}
	var num_list []int
	for _, idx := range indexlist {
		index := idx["idx"]
		index1, _ := strconv.Atoi(index)
		num_list = append(num_list, index1)
	}
	Shuffle(num_list)
	var randsongs []map[string]string
	for _, f := range num_list {
		ff := strconv.Itoa(f)
		log.Printf("this is f: %s", f)
		filter := bson.D{{"idx", ff}}
		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
		defer ampgosetup.Close(client, ctx, cancel)
		ampgosetup.CheckError(err, "MongoDB connection has failed")
		collection := client.Database("maindb").Collection("maindb")
		var rpics map[string]string
		err = collection.FindOne(context.Background(), filter).Decode(&rpics)
		if err != nil { log.Fatal(err) }
		log.Printf("rpics: %v", rpics)

		randsongs = append(randsongs, rpics)
	}
	log.Println(randsongs)
	var plz AmpgoRandomPlaylistData
	plz.PlayListName = plname
	plz.PlayListID = plID
	plz.PlayList = randsongs[:plcount]
	log.Println("This is plz")
	log.Println(plz)
	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
	ServerCheckError(err3, "Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", &plz)
	ServerCheckError(err2, "plz insertion has failed")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Playlist Created")
}

func allPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("randplaylists").Collection("randplaylists")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "allIdx has failed")
	var allplaylists []AmpgoRandomPlaylistData
	if err = cur.All(context.TODO(), &allplaylists); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allplaylists)
}

// func setCurrentPlaylist(w http.ResponseWriter, r *http.Request) {

// }

// func getCurrentPlaylist(w http.ResponseWriter, r *http.Request) {

// }

func init() {
	ampgosetup.SetUpCheck()
	var logging_status string = StartServerLogging()
	log.Println(logging_status)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/SetUp", setUpHandler)
	r.HandleFunc("/Home", homeHandler)

	r.HandleFunc("/InitArtistInfo", initArtistInfoHandler)
	// r.HandleFunc("/AlbumsForArtist", albumsForArtistHandler)
	r.HandleFunc("/AlbumsForArtist2", albumsForArtist2Handler)

	r.HandleFunc("/SongsForAlbum", songsForAlbumHandler)

	r.HandleFunc("/AddPlaylist", addPlaylistHandler)
	r.HandleFunc("/CreateRandomPlaylist", createRandomPlaylistHandler)
	r.HandleFunc("/AllPlaylists", allPlaylistsHandler)
	
	r.HandleFunc("/RandomPics", randomPicsHandler)
	////////////////////////////////////////////////////////////////

	// r.HandleFunc("/SetCurrentPlaylist", setCurrentPlaylistHandler) 
	// r.HandleFunc("/GetCurrentPlaylist", getCurrentPlaylistHandler) 
	

	// r.HandleFunc("/DeletPlaylist", deletePlaylistHandler)
	// r.HandleFunc("/EditPlaylist", editPlaylistHandler)
	// r.HandleFunc("/AddSongToPlaylist", addSongToPlaylistHandler)
	// r.HandleFunc("/DeleteSongFromPlaylist", deleteSongFromPlaylistHandler)

	
	/////////////////////////////////////////////////////

	r.HandleFunc("/InitAlbumInfo", initalbumInfoHandler)
	r.HandleFunc("/InitialSongInfo", initialsongInfoHandler)
	


	

	// r.HandleFunc("/ArtistAlpha", artistPageHandler)
	// r.HandleFunc("/AlbumAlpha", albumPageHandler)
	// r.HandleFunc("/TitleAlpha", titlePageHandler)

	// r.HandleFunc("/ArtistInfo", artistInfoHandler)
	// r.HandleFunc("/AlbumInfo", albumInfoHandler)
	// r.HandleFunc("/SongInfo", songInfoHandler)




	// r.HandleFunc("/ImageSongsForAlbum", imageSongsForAlbumHandler)
	// r.HandleFunc("/SongsForAlbum", songsForAlbumHandler)


	

	
	// r.HandleFunc("/RamdomAlbumPicPlaySong", ramdomAlbumPicPlaySongHandler)
	// r.HandleFunc("/Stats", statsHandler)
	// r.HandleFunc("/PathArt", pathArtHandler)
	// r.HandleFunc("/SongSearch", songSearchHandler)
	// r.HandleFunc("/AlbumSearch", albumSearchHandler)
	// r.HandleFunc("/ArtistSearch", artistSearchHandler)
	// r.HandleFunc("/AllPlaylists", allPlaylistsHandler)
	// r.HandleFunc("/AddPlayListNameToDB", addPlayListNameToDBHandler)
	// r.HandleFunc("/AddSongsToPlistDB", addSongsToPlistDBHandler)
	// r.HandleFunc("/AllPlaylistSongsFromDB", allPlaylistSongsFromDBHandler)
	// r.HandleFunc("/CreatePlayerPlaylist", createPlayerPlaylistHandler)
	// r.HandleFunc("/AddRandomPlaylist", addRandomPlaylistHandler)
	// r.HandleFunc("/DeletePlaylistFromDB", deletePlaylistFromDBHandler)
	// r.HandleFunc("/DeleteSongFromPlaylist", deleteSongFromPlaylistHandler)


	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/root/static/"))))
	r.PathPrefix("/fsData/").Handler(http.StripPrefix("/fsData/", http.FileServer(http.Dir("/root/fsData/"))))

	// s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	// r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/static/"))))
	http.ListenAndServe(":9090", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), 
		handlers.AllowedOrigins([]string{"*"}))(r))


}


func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
    collection := client.Database(dataBase).Collection(col)
    result, err := collection.InsertOne(ctx, doc)
    return result, err
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}

func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = 0x80
	uuid[4] = 0x40
	boo := hex.EncodeToString(uuid)
	return boo, nil
}