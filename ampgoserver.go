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
	"sort"
	"net/http"
	// "html/template"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/bson"
	"context"
    // "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/globalsign/mgo"
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
	HSImage string              `bson:"hsimage"`
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

var OFFSET string = os.Getenv("AMPGO_OFFSET")

func StartServerLogging() string {
	var logtxtfile string = os.Getenv("AMPGO_SERVER_LOG_PATH")
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logtxtfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return "Logging started"
}


func ServerCheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Println(msg)
		panic(err)
	}
}

func sfdbCon() *mgo.Session {
	s, err := mgo.Dial("mongodb://db:27017/ampgodb")
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}

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

func initialArtistInfoHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("artistview").Collection("artistview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initialArtistInfo find has failed")
	var av []ArtVIEW
	if err = cur.All(context.TODO(), &av); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s this is av", av)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&av)
	log.Println("Initial Artist Info Complete")
}

func initialalbumInfoHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initialalbumInfo find has failed")
	var albv []AlbvieW
	if err = cur.All(context.TODO(), &albv); err != nil {
	}
	log.Printf("%s this is albv", albv)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&albv)
	log.Println("Initial Album Info Complete")
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

func artistPageHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	opts := options.Distinct()
	opts.SetMaxTime(2 * time.Second)
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("artistview").Collection("artistview")
	DD1, err2 := collection.Distinct(context.TODO(), "page", filter, opts)
	ServerCheckError(err2, "MongoDB distinct album has failed")
	var ARDist []string
	for _, DD := range DD1 {
		zoo := fmt.Sprintf("%s", DD)
		ARDist = append(ARDist, zoo)
	}
	sort.Strings(ARDist)
	log.Println("ArtistAlpha is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ARDist)
}

func albumPageHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	opts := options.Distinct()
	opts.SetMaxTime(2 * time.Second)
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("albumview").Collection("albumview")
	DD1, err2 := collection.Distinct(context.TODO(), "albumpage", filter, opts)
	ServerCheckError(err2, "MongoDB distinct album has failed")
	var ALDist []string
	for _, DD := range DD1 {
		zoo := fmt.Sprintf("%s", DD)
		ALDist = append(ALDist, zoo)
	}
	sort.Strings(ALDist)
	log.Println("AlbumAlpha is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ALDist)
}

func titlePageHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	opts := options.Distinct()
	opts.SetMaxTime(2 * time.Second)
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("maindb").Collection("maindb")
	DD1, err2 := collection.Distinct(context.TODO(), "titlepage", filter, opts)
	ServerCheckError(err2, "MongoDB distinct album has failed")
	var TDist []string
	for _, DD := range DD1 {
		zoo := fmt.Sprintf("%s", DD)
		TDist = append(TDist, zoo)
	}
	sort.Strings(TDist)
	log.Println("TitleAlpha is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TDist)
}

func songInfoHandler(w http.ResponseWriter, r *http.Request) {
	pagenum := r.URL.Query().Get("selected")

	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{"titlepage", pagenum}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "ArtPipeline find has failed")
	var SIS []map[string]string
	if err = cur.All(context.TODO(), &SIS); err != nil {
		log.Fatal(err)
	}
	// ses := sfdbCon()
	// defer ses.Close()
	// MAINc := ses.DB("maindb").C("maindb")
	// b1 := bson.M{"page": pagenum}
	// b2 := bson.M{"_id": 0, "title": 1, "fileID": 1, "artist": 1}
	// var SIS []map[string]string
	// err := MAINc.Find(b1).Select(b2).All(&SIS)
	// if err != nil {
	// 	log.Println("song info has fucked up")
	// 	log.Println(err)
	// }
	log.Println("SongInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SIS)
}

func albumInfoHandler(w http.ResponseWriter, r *http.Request) {
	pagenum := r.URL.Query().Get("selected")
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{"albumpage", pagenum}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
	opts.SetProjection(b2)
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "AlbPipeline find has failed")
	var AI []AlbvieW
	if err = cur.All(context.TODO(), &AI); err != nil {
		log.Fatal(err)
	}

	// ses := sfdbCon()
	// defer ses.Close()
	// ALBVc := ses.DB("albview").C("albview")
	// b1 := bson.M{"page": pagenum}
	// b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
	// var AI []AlbvieW
	// err := ALBVc.Find(b1).Select(b2).All(&AI)
	// if err != nil {
	// 	log.Println("AlbumInfo has fucked up")
	// 	log.Println(err)
	// }
	log.Println("AlbumInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&AI)
}

func artistInfoHandler(w http.ResponseWriter, r *http.Request) {
	pagenum := r.URL.Query().Get("selected")
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "convert to int64 has failed")
	filter := bson.D{{"page", pagenum}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
	opts.SetProjection(b2)
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("artistview").Collection("artistview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "ArtPipeline find has failed")
	var ARTI []ArtVIEW
	if err = cur.All(context.TODO(), &ARTI); err != nil {
		log.Fatal(err)
	}
	log.Println("ArtistInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ARTI)
}

func imageSongsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
	albumid := r.URL.Query().Get("selected")
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	ServerCheckError(err, "ParseInt has failed")
	filter := bson.D{{"albumID", albumid}}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	b2 := bson.M{"_id": 0, "album": 1, "songs": 1, "picPath": 1}
	opts.SetProjection(b2)
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "imageSongsForAlbumHandler has failed")
	var iM []iMgfa
	if err = cur.All(context.TODO(), &iM); err != nil {
		log.Fatal(err)
	}
	// ses := sfdbCon()
	// defer ses.Close()
	// ALBc := ses.DB("albview").C("albview")
	// b1 := bson.M{"albumID": albid}
	// b2 := bson.M{"_id": 0, "album": 1, "songs": 1, "picPath": 1}
	// var iM []iMgfa
	// err := ALBc.Find(b1).Select(b2).One(&iM)
	// if err != nil {
	// 	log.Println("gimage song for album fucked up")
	// 	log.Println(err)
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&iM)
}

func randomPicsHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	limit, err := strconv.ParseInt(OFFSET, 10, 64)
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetProjection(bson.M{"_id": 0, "idx": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "imageSongsForAlbumHandler has failed")
	var indexliststring []map[string]string
	if err = cur.All(context.TODO(), &indexliststring); err != nil {
		log.Fatal(err)
	}
	var indexlistint []int 
	for _, idx := range indexliststring {
		idxint, err := strconv.Atoi(idx["idx"])
		ServerCheckError(err, "ParseInt has failed")
		indexlistint = append(indexlistint, idxint)
	}
	log.Println(indexlistint)
	var albumcount int = len(indexlistint) - 1
	log.Println(albumcount)
	log.Printf("%T albumcount", albumcount)
	// ses := sfdbCon()
	// defer ses.Close()
	// ALBc := ses.DB("coverart").C("coverart")
	// albumcount, err := ALBc.Count()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	var min int = 1
	maxx := albumcount
	max := int(maxx)

	var five_rand_num []string
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		random11 := rand.Intn(max - min) + min
		random1 := strconv.Itoa(random11)
		time.Sleep(50 * time.Millisecond)
		five_rand_num = append(five_rand_num, random1)
	}
	log.Println(five_rand_num)

	var randpics []map[string]string
	for _, ff := range five_rand_num {
		f, err := strconv.Itoa(ff)
		ServerCheckError(err, "strconv.Atoi has failed")
		filter := bson.D{{"index", f}}
		limit, err := strconv.ParseInt(OFFSET, 10, 64)
		ServerCheckError(err, "Int conversion has failed")
		opts := options.Find()
		opts.SetLimit(int64(limit))
		opts.SetProjection(bson.M{"_id": 0})
		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
		defer ampgosetup.Close(client, ctx, cancel)
		ServerCheckError(err, "MongoDB connection has failed")
		coll := client.Database("coverart").Collection("coverart")
		cur, err := coll.Find(context.TODO(), filter, opts)
		ServerCheckError(err, "randomPicsHandler find has failed")
		var iM map[string]string
		if err = cur.All(context.TODO(), &iM); err != nil {
			log.Fatal(err)
		}
		randpics = append(randpics, iM)
	}
	// 	// ses := sfdbCon()
	// 	// defer ses.Close()
	// 	// ALBc := ses.DB("coverart").C("coverart")
	// 	// b1 := bson.M{"index": f}
	// 	// b2 := bson.M{"_id": 0}
	// 	// var iM map[string]string
	// 	// err := ALBc.Find(b1).Select(b2).One(&iM)
	// 	// if err != nil {
	// 	// 	log.Println("gimage song for album fucked up")
	// 	// 	log.Println(err)
	// 	// }
		
	// 	// return randpics
		
	
	// fmt.Println(randpics)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(randpics)

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

// func allPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	plc := ampgolib.PlaylistCheck()
// 	if plc != 1 {
// 		json.NewEncoder(w).Encode("Please create a playlist")
// 	} else {
// 		allpls := ampgolib.AllPlayLists()
// 		json.NewEncoder(w).Encode(allpls)
// 	}
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

// func addRandomPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plname := r.URL.Query().Get("playlistname")
// 	plcount := r.URL.Query().Get("playlistcount")
// 	rpl := ampgolib.AddRandomPlaylist(plname, plcount)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(rpl)
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



func init() {
	ampgosetup.SetUpCheck()
	var logging_status string = StartServerLogging()
	log.Println(logging_status)
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/SetUp", setUpHandler)
	r.HandleFunc("/Home", homeHandler)

	r.HandleFunc("/InitialArtistInfo", initialArtistInfoHandler)
	r.HandleFunc("/InitialAlbumInfo", initialalbumInfoHandler)
	r.HandleFunc("/InitialSongInfo", initialsongInfoHandler)

	r.HandleFunc("/ArtistAlpha", artistPageHandler)
	r.HandleFunc("/AlbumAlpha", albumPageHandler)
	r.HandleFunc("/TitleAlpha", titlePageHandler)

	r.HandleFunc("/ArtistInfo", artistInfoHandler)
	r.HandleFunc("/AlbumInfo", albumInfoHandler)
	r.HandleFunc("/SongInfo", songInfoHandler)




	r.HandleFunc("/ImageSongsForAlbum", imageSongsForAlbumHandler)
	r.HandleFunc("/RandomPics", randomPicsHandler)


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

	
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/static/"))))
	http.ListenAndServe(":9090", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), 
		handlers.AllowedOrigins([]string{"*"}))(r))


}


// func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
// 	defer cancel()

// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

// func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
 
//     ctx, cancel := context.WithTimeout(context.Background(),
//                                        30 * time.Second)
//     client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
//     return client, ctx, cancel, err
// }

// func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
 
//     collection := client.Database(dataBase).Collection(col)
     
//     result, err := collection.InsertOne(ctx, doc)
//     return result, err
// }