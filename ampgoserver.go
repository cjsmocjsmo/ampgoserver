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
	"fmt"
<<<<<<< HEAD
	"log"
	"math/rand"
	"os"
	// "os/exec"
	// "io"
	// "strconv"
	"time"
	// "path/filepath"
	// "bufio"
=======
	"os"
	"sort"

	// "log"
	"math/rand"
	"strconv"
	"time"
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68

	// "path"
	// "strings"
	// "io/ioutil"
	// "sort"
	"context"
	"encoding/hex"
<<<<<<< HEAD
	// "encoding/json"
=======
	"encoding/json"
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68
	"net/http"

	"github.com/cjsmocjsmo/ampgosetup"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
<<<<<<< HEAD
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "crypto/tls"
	// "golang.org/x/crypto/acme/autocert"
=======
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68
)

// type plist struct {
// 	PLName string              `bson:"PLName"`
// 	PLId   string              `bson:"PLId"`
// 	Songs  []map[string]string `bson:"Songs"`
// }

// type iMgfa struct {
// 	Album   string              `bson:"album"`
// 	PicPath string              `bson:"picPath"`
// 	Songs   []map[string]string `bson:"songs"`
// }

// type rAlbinfo struct {
// 	Songs   []map[string]string `bson:"songs"`
// 	HSImage string              `bson:"hsimage"`
// }

// type voodoo struct {
// 	Playlists []map[string]string `bson:"playlists"`
// }

//ArtVIEW exported
type ArtVIEW struct {
	Artist   string              `bson:"artist"`
	ArtistID string              `bson:"artistID"`
	Albums   []map[string]string `bson:"albums"`
	Page     string              `bson:"page"`
	Idx      string              `bson:"idx"`
}

type AlbVieW2 struct {
	Artist      string              `bson:"artist"`
	ArtistID    string              `bson:"artistID"`
	Album       string              `bson:"album"`
	AlbumID     string              `bson:"albumID"`
	Songs       []map[string]string `bson:"songs"`
	AlbumPage   string              `bson:"albumpage"`
	NumSongs    string              `bson:"numsongs"`
	PicPath     string              `bson:"picPath"`
	Idx         string              `bson:"idx"`
	PicHttpAddr string              `bson:"picHttpAddr"`
}

type AmpgoRandomPlaylistData struct {
	PlayListName  string              `bson:"playlistname"`
	PlayListID    string              `bson:"playlistID"`
	PlayListCount string              `bson:"playlistcount"`
	PlayList      []map[string]string `bson:"playlist"`
}

var OFFSET string = os.Getenv("AMPGO_OFFSET")

<<<<<<< HEAD
func ServerCheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Println(msg)
		panic(err)
	}
}

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
=======
// func RemoveLogFile(logtxtfile string) bool {
// 	// var logtxtfile string = os.Getenv("AMPGO_SERVER_LOG_PATH")
// 	var result bool
// 	_, err := os.Stat(logtxtfile)
//     if err == nil {
//         fmt.Printf("logfile %s exists removing", logtxtfile)
// 		os.Remove(logtxtfile)
// 		result = true
//     } else if os.IsNotExist(err) {
//         fmt.Printf("logfile %s does not exists", logtxtfile)
// 		result = true
//     } else {
//         fmt.Printf("logfile %s stat error: %v", logtxtfile, err)
// 		result = false
//     }
// 	return result
// }
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68

// func StartServerLogging() string {
// 	var logtxtfile string = os.Getenv("AMPGO_SERVER_LOG_PATH")
// 	result := RemoveLogFile(logtxtfile)
// 	if result {
// 		file, err := os.OpenFile(logtxtfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.SetOutput(file)
// 	} else {
// 		fmt.Println("Unable to setup logging")
// 	}
// 	return "Logging started"
// }

<<<<<<< HEAD

func init() {
	fmt.Println("Starting init")
	log.Println("Starting init")
	StartServerLogging()
	ampgosetup.Setup()


}



// func setUpHandler(w http.ResponseWriter, r *http.Request) {
// 	ampgosetup.Setup()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Setup Complete")
// 	//this needs work
// 	log.Println("Setup Complete")
// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<html><h1> Hello From Ampgo Home \n It works </h1></html>")
	log.Println("homeHandler is complete")
=======
func ServerCheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(msg)
		panic(err)
	}
}

func setUpHandler(w http.ResponseWriter, r *http.Request) {
	ampgosetup.Setup()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Setup Complete")
	//this needs work
	fmt.Println("Setup Complete")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello From Ampgo Home \n It works")
	fmt.Println("homeHandler is complete")
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68
}

///////////////////////////////////////////////////////////////////////////////
//Artist Stuff
///////////////////////////////////////////////////////////////////////////////

<<<<<<< HEAD
// func artistInfoByPageHandler(w http.ResponseWriter, r *http.Request) {
// 	page := r.URL.Query().Get("page")
// 	filter := bson.M{"page": page}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("artistview").Collection("artistview")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "initArtistInfo find has failed")
// 	var allartist []ArtVIEW
// 	if err = cur.All(context.TODO(), &allartist); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is allartist-", allartist)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allartist)
// 	log.Println("Init Artist Info Complete")
// }

// func ArtViewFindOne(db string, coll string, filtertype string, filterstring string) ArtVIEW {
// 	filter := bson.M{filtertype: filterstring}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database(db).Collection(coll)
// 	var results ArtVIEW
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return results
// }

// func albumsForArtist2Handler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Starting albumsForArtistHandler")
// 	var artistid string = r.URL.Query().Get("selected")
// 	log.Printf("%s this is artistid", artistid)
// 	log.Printf("%T this is artistid type", artistid)
// 	allalbums := ArtViewFindOne("artistview", "artistview", "artistID", artistid)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allalbums.Albums)
// 	log.Println("Init Artist Info Complete")
// }

// func albumsForArtistHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Starting albumsForArtistHandler")
// 	var artistid string = r.URL.Query().Get("selected")
// 	log.Printf("%s this is artistid", artistid)
// 	log.Printf("%T this is artistid type", artistid)
// 	filter := bson.M{"artistID": artistid}
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

// func songsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Starting songsForAlbumHandler")
// 	var albumid string = r.URL.Query().Get("selected")
// 	log.Printf("%s this is albumid", albumid)
// 	log.Printf("%T this is albumid type", albumid)
// 	filter := bson.M{"albumID": albumid}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("maindb")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "songsForAlbumHandler find has failed")
// 	var allalbum []map[string]string
// 	if err = cur.All(context.TODO(), &allalbum); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is allalbum-", allalbum)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allalbum)
// 	log.Println("songsForAlbumHandler Complete")
// }
=======
func artistInfoByPageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	filter := bson.M{"page": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("artistview").Collection("artistview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initArtistInfo find has failed")
	var allartist []ArtVIEW
	if err = cur.All(context.TODO(), &allartist); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is allartist-", allartist)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allartist)
	fmt.Println("Init Artist Info Complete")
}

func ArtViewFindOne(db string, coll string, filtertype string, filterstring string) ArtVIEW {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results ArtVIEW
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}

func albumsForArtist2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting albumsForArtistHandler")
	var artistid string = r.URL.Query().Get("selected")
	fmt.Printf("%s this is artistid", artistid)
	fmt.Printf("%T this is artistid type", artistid)
	allalbums := ArtViewFindOne("artistview", "artistview", "artistID", artistid)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbums.Albums)
	fmt.Println("Init Artist Info Complete")
}

func albumsForArtistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting albumsForArtistHandler")
	var artistid string = r.URL.Query().Get("selected")
	fmt.Printf("%s this is artistid", artistid)
	fmt.Printf("%T this is artistid type", artistid)
	filter := bson.M{"artistID": artistid}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "songs": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initArtistInfo find has failed")
	var allalbum []map[string]string
	if err = cur.All(context.TODO(), &allalbum); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is allalbum-", allalbum)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbum)
	fmt.Println("Init Album Info Complete")
}

func songsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting songsForAlbumHandler")
	var albumid string = r.URL.Query().Get("selected")
	fmt.Printf("%s this is albumid", albumid)
	fmt.Printf("%T this is albumid type", albumid)
	filter := bson.M{"albumID": albumid}
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
		fmt.Println(err)
	}
	fmt.Printf("%s this is allalbum-", allalbum)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbum)
	fmt.Println("songsForAlbumHandler Complete")
}
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68

// ///////////////////////////////////////////////////////////////////////////////
// //Album Stuff
// ///////////////////////////////////////////////////////////////////////////////

<<<<<<< HEAD
// func albumInfoByPageHandler(w http.ResponseWriter, r *http.Request) {
// 	page := r.URL.Query().Get("page")
// 	filter := bson.M{"albumpage": page}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("albumview").Collection("albumview")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "initAlbumInfo find has failed")
// 	var allalbums []AlbVieW2
// 	if err = cur.All(context.TODO(), &allalbums); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is allalbums", allalbums)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allalbums)
// 	log.Println("Init albumsInfo Complete")
// }

// func songInfoByPageHandler(w http.ResponseWriter, r *http.Request) {
// 	page := r.URL.Query().Get("page")
// 	filter := bson.M{"titlepage": page}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1, "httpaddr": 1})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("maindb")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "ArtPipeline find has failed")
// 	var tv []map[string]string
// 	if err = cur.All(context.TODO(), &tv); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is tv", tv)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&tv)
// 	log.Println("Initial Song Info Complete")
// }

// func playSongHandler(w http.ResponseWriter, r *http.Request) {
// 	fileid := r.URL.Query().Get("selected")
// 	filter := bson.M{"fileID": fileid}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("maindb").Collection("maindb")
// 	var results map[string]string
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var newresults = map[string]string{
// 		"httpaddr":    results["httpaddr"],
// 		"artist":      results["artist"],
// 		"album":       results["album"],
// 		"title":       results["title"],
// 		"duration":    results["duration"],
// 		"picHttpAddr": results["picHttpAddr"],
// 	}
// 	log.Printf("%s this is playSong", newresults)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&newresults)
// 	log.Println("playSong Song Info Complete")
// }

// func playPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	playlistid := r.URL.Query().Get("selected")
// 	filter := bson.M{"playlistID": playlistid}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("playlistdb").Collection("playlistdb")
// 	var results map[string]string
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is playPlaylist", results)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&results)
// 	log.Println("playPlaylist Info Complete")
// }

// func randomPicsHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.D{{}}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0, "index": 1})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("coverart").Collection("coverart")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "randomPicsHandler has failed")
// 	var indexliststring []map[string]string
// 	if err = cur.All(context.TODO(), &indexliststring); err != nil {
// 		log.Fatal(err)
// 	}
// 	var num_list []int
// 	for _, idx := range indexliststring {
// 		indexx := idx["index"]
// 		index1, _ := strconv.Atoi(indexx)
// 		num_list = append(num_list, index1)
// 	}
// 	shuffle(num_list)
// 	log.Println(num_list)
// 	var randpics []string
// 	for _, f := range num_list[:12] {
// 		log.Printf("f type: %T", f)
// 		ff := strconv.Itoa(f)
// 		log.Printf("ff type %T", ff)
// 		log.Printf("ff type %s", ff)
// 		filter := bson.M{"index": ff}
// 		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 		defer ampgosetup.Close(client, ctx, cancel)
// 		ampgosetup.CheckError(err, "MongoDB connection has failed")
// 		collection := client.Database("coverart").Collection("coverart")
// 		var rpics map[string]string
// 		err = collection.FindOne(context.Background(), filter).Decode(&rpics)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		randpics = append(randpics, rpics["imagehttpaddr"])
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(randpics)
// }

// ///////////////////////////////////////////////////////////////////////////////
// //Playlist Stuff
// ///////////////////////////////////////////////////////////////////////////////
// func AllPlayListsFunc() []AmpgoRandomPlaylistData {
// 	filter := bson.M{}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("randplaylists").Collection("randplaylists")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "allIdx has failed")
// 	var allplaylists []AmpgoRandomPlaylistData
// 	if err = cur.All(context.TODO(), &allplaylists); err != nil {
// 		log.Fatal(err)
// 	}
// 	return allplaylists
// }

// func allPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
// 	allplaylists := AllPlayListsFunc()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allplaylists)
// }

// func deletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plid := r.URL.Query().Get("playlistid")
// 	filter := bson.M{"playlistID": plid}
// 	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
// 	ServerCheckError(err3, "Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := DeleteOne(client, ctx, "randplaylists", "randplaylists", filter)
// 	ServerCheckError(err2, "deleteplaylist has failed")
// 	allplaylists := AllPlayListsFunc()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allplaylists)
// }

// func addPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plname := r.URL.Query().Get("name")
// 	plID, _ := UUID()
// 	var emptymap []map[string]string
// 	var emptyitem map[string]string = map[string]string{"title": "None Found"}
// 	emptymap = append(emptymap, emptyitem)
// 	var plzz AmpgoRandomPlaylistData
// 	plzz.PlayListName = plname
// 	plzz.PlayListID = plID
// 	plzz.PlayList = emptymap
// 	log.Println("This is plzz")
// 	log.Println(plzz)
// 	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
// 	ServerCheckError(err3, "Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", &plzz)
// 	ServerCheckError(err2, "plz insertion has failed")
// 	allplaylists := AllPlayListsFunc()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allplaylists)
// }

// func shuffle(slice []int) {
// 	r := rand.New(rand.NewSource(time.Now().Unix()))
// 	for n := len(slice); n > 0; n-- {
// 		randIndex := r.Intn(n)
// 		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
// 	}
// }

// // test with curl http://192.168.0.91:9090/CreateRandomPlaylist?songcount=25&&name=RucRandom
// func genrandom(maxx int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(maxx)
// }

// func addRandomPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	plc := r.URL.Query().Get("songcount")
// 	plname := r.URL.Query().Get("name")
// 	log.Printf("planame: %s", plname)
// 	log.Printf("plc: %s", plc)
// 	plcount, _ := strconv.Atoi(plc)
// 	plID, _ := UUID()
// 	filter := bson.D{{}}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("songtotal").Collection("total")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "allIdx has failed")
// 	var num_map []map[string]string
// 	if err = cur.All(context.TODO(), &num_map); err != nil {
// 		log.Println("randplaylist dbcall has fucked up")
// 		log.Fatal(err)
// 	}
// 	log.Println(num_map)
// 	somenum := ""
// 	for _, item := range num_map {
// 		somenum = item["total"]
// 	}
// 	log.Println(somenum)
// 	var num_list []int
// 	log.Println(plcount)
// 	for i := 1; i <= plcount; i++ {
// 		newTotal, _ := strconv.Atoi(somenum)
// 		log.Println(newTotal)
// 		ranN := genrandom(newTotal)
// 		log.Println(ranN)
// 		num_list = append(num_list, ranN)
// 	}
// 	log.Println(num_list)
// 	shuffle(num_list)
// 	var randsongs []map[string]string
// 	for _, f := range num_list {
// 		if len(randsongs) == plcount {
// 			break
// 		}
// 		ff := strconv.Itoa(f)
// 		log.Println(ff)
// 		log.Printf("this is ff type: %T", ff)
// 		log.Printf("this is f type: %T", f)
// 		filter := bson.M{"idx": ff}
// 		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 		defer ampgosetup.Close(client, ctx, cancel)
// 		ampgosetup.CheckError(err, "MongoDB connection has failed")
// 		collection := client.Database("maindb").Collection("maindb")
// 		var rplaylists map[string]string
// 		err = collection.FindOne(context.Background(), filter).Decode(&rplaylists)
// 		if err != nil {
// 			log.Println(ff)
// 			continue
// 		} //log.Fatal(err) }
// 		randsongs = append(randsongs, rplaylists)
// 	}
// 	var plz AmpgoRandomPlaylistData
// 	plz.PlayListName = plname
// 	plz.PlayListID = plID
// 	plz.PlayListCount = strconv.Itoa(len(randsongs))
// 	plz.PlayList = randsongs
// 	log.Println("This is plz")
// 	log.Println(plz)
// 	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
// 	ServerCheckError(err3, "Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", &plz)
// 	ServerCheckError(err2, "plz insertion has failed")
// 	allplaylists := AllPlayListsFunc()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allplaylists)
// }

// func getCurrentPlayListNameHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.M{"record": "1"}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "updateCurrentPlayListName: MongoDB connection has failed")
// 	collection := client.Database("curplaylistname").Collection("curplaylistname")
// 	var results map[string]string
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&results)
// }

// func updateCurrentPlayListNameHandler(w http.ResponseWriter, r *http.Request) {
// 	curplaylistname := r.URL.Query().Get("curplaylistname")
// 	curplaylistid := r.URL.Query().Get("curplaylistid")
// 	filter := bson.M{"record": "1"}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "updateCurrentPlayListName: MongoDB connection has failed")
// 	collection := client.Database("curplaylistname").Collection("curplaylistname")
// 	var results map[string]string
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	update := bson.M{"$set": bson.M{"curplaylistname": curplaylistname, "curplaylistID": curplaylistid}}
// 	_, err = collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var find_results map[string]string
// 	err = collection.FindOne(context.Background(), filter).Decode(&find_results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&find_results)
// }

// func playListByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	playlistid := r.URL.Query().Get("playlistid")
// 	filter := bson.M{"playlistID": playlistid}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "updateCurrentPlayListName: MongoDB connection has failed")
// 	collection := client.Database("randplaylists").Collection("randplaylists")
// 	var results AmpgoRandomPlaylistData
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&results)
// }

// func songInfoFindOne(db string, coll string, filtertype string, filterstring string) map[string]string {
// 	filter := bson.M{filtertype: filterstring}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database(db).Collection(coll)
// 	var results map[string]string
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return results
// }

// func playlistInfoFromPlaylistID(db string, coll string, filtertype string, filterstring string) AmpgoRandomPlaylistData {
// 	filter := bson.M{filtertype: filterstring}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database(db).Collection(coll)
// 	var results AmpgoRandomPlaylistData
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return results
// }

// func increasePlayListCount(astring string) string {
// 	newint, _ := strconv.Atoi(astring)
// 	newnewint := newint + 1
// 	return strconv.Itoa(newnewint)
// }

// func decreasePlayListCount(astring string) string {
// 	newint, _ := strconv.Atoi(astring)
// 	newnewint := newint - 1
// 	return strconv.Itoa(newnewint)
// }

// func addSongToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	fileID := r.URL.Query().Get("fileid")
// 	plid := r.URL.Query().Get("playlistid")
// 	log.Printf("fileID: %s", fileID)
// 	log.Printf("plid: %s", plid)

// 	songinfo := songInfoFindOne("maindb", "maindb", "fileID", fileID)
// 	log.Println("This is songinfo")
// 	log.Println(songinfo)

// 	playlistInfo := playlistInfoFromPlaylistID("randplaylists", "randplaylists", "playlistID", plid)
// 	log.Println("this is playlistinfo")
// 	log.Println(playlistInfo)

// 	playlistInfo.PlayList = append(playlistInfo.PlayList, songinfo)
// 	newcount := increasePlayListCount(playlistInfo.PlayListCount)

// 	update := bson.M{"$set": bson.M{
// 		"playlistcount": newcount,
// 		"playlist":      playlistInfo.PlayList,
// 	},
// 	}

// 	filter := bson.M{"playlistID": plid}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("randplaylists").Collection("randplaylists")
// 	_, err = collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Playlist updated")
// }

// func deleteSongFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
// 	fileID := r.URL.Query().Get("fileid")
// 	plid := r.URL.Query().Get("playlistid")
// 	log.Printf("fileID: %s", fileID)
// 	log.Printf("plid: %s", plid)

// 	PlayListInfo := playlistInfoFromPlaylistID("randplaylists", "randplaylists", "playlistID", plid)
// 	var newMap []map[string]string
// 	for _, song := range PlayListInfo.PlayList {
// 		if song["fileID"] == fileID {
// 			continue
// 		} else {
// 			newMap = append(newMap, song)
// 		}
// 	}

// 	newcount := decreasePlayListCount(PlayListInfo.PlayListCount)

// 	log.Println(newMap)
// 	log.Println(newcount)

// 	update := bson.M{"$set": bson.M{
// 		"playlistcount": newcount,
// 		"playlist":      newMap,
// 	},
// 	}

// 	filter := bson.M{"playlistID": plid}
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "MongoDB connection has failed")
// 	collection := client.Database("randplaylists").Collection("randplaylists")
// 	_, err = collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	allplaylists := AllPlayListsFunc()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allplaylists)

// }

// ///////////////////////////////////////////////////////////////////////////////
// // alphabet stuff
// ///////////////////////////////////////////////////////////////////////////////

// func artistAlphaHandler(w http.ResponseWriter, r *http.Request) {
// 	alpha := r.URL.Query().Get("alpha")
// 	filter := bson.M{}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "artistAlpha: MongoDB connection has failed")
// 	coll := client.Database("artistalpha").Collection(alpha)
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "artistAlpha: allIdx has failed")
// 	var allItems []map[string]string
// 	if err = cur.All(context.TODO(), &allItems); err != nil {
// 		log.Fatal(err)
// 	}

// 	if len(allItems) < 1 {
// 		z := []map[string]string{}
// 		x := map[string]string{}
// 		z = append(z, x)
// 		var noresult ArtVIEW = ArtVIEW{
// 			Artist:   "None Found",
// 			ArtistID: "None Found",
// 			Albums:   z,
// 			Page:     "None Found",
// 			Idx:      "None Found",
// 		}
// 		zoo := []ArtVIEW{}
// 		zoo = append(zoo, noresult)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(&zoo)
// 	} else {
// 		var NewAllItems []ArtVIEW
// 		for _, item := range allItems {
// 			filter := bson.M{"artist": item["artist"]}
// 			client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 			defer ampgosetup.Close(client, ctx, cancel)
// 			ServerCheckError(err, "artistAlpha: MongoDB connection has failed")
// 			collection := client.Database("artistview").Collection("artistview")
// 			var results ArtVIEW
// 			err = collection.FindOne(context.Background(), filter).Decode(&results)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			NewAllItems = append(NewAllItems, results)
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(&NewAllItems)
// 	}

// }

// func albumAlphaHandler(w http.ResponseWriter, r *http.Request) {
// 	alpha := r.URL.Query().Get("alpha")
// 	filter := bson.D{{}}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "albumAlpha: MongoDB connection has failed")
// 	coll := client.Database("albumalpha").Collection(alpha)
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "albumAlpha: allIdx has failed")
// 	var allItems []map[string]string
// 	if err = cur.All(context.TODO(), &allItems); err != nil {
// 		log.Fatal(err)
// 	}

// 	if len(allItems) < 1 {
// 		z := []map[string]string{}
// 		x := map[string]string{}
// 		z = append(z, x)
// 		var noresult AlbVieW2 = AlbVieW2{
// 			Artist:      "None Found",
// 			ArtistID:    "None Found",
// 			Album:       "None Found",
// 			AlbumID:     "None Found",
// 			Songs:       z,
// 			AlbumPage:   "None Found",
// 			NumSongs:    "None Found",
// 			PicPath:     "None Found",
// 			Idx:         "None Found",
// 			PicHttpAddr: "None Found",
// 		}
// 		zoo := []AlbVieW2{}
// 		zoo = append(zoo, noresult)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(&zoo)
// 	} else {
// 		var NewAllItems []AlbVieW2
// 		for _, item := range allItems {
// 			filter := bson.M{"album": item["album"]}
// 			client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 			defer ampgosetup.Close(client, ctx, cancel)
// 			ServerCheckError(err, "MongoDB connection has failed")
// 			collection := client.Database("albumview").Collection("albumview")
// 			var results AlbVieW2
// 			err = collection.FindOne(context.Background(), filter).Decode(&results)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			NewAllItems = append(NewAllItems, results)
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(&NewAllItems)
// 	}
// }

// func songAlphaHandler(w http.ResponseWriter, r *http.Request) {
// 	alpha := r.URL.Query().Get("alpha")
// 	filter := bson.D{{}}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
// 	defer ampgosetup.Close(client, ctx, cancel)
// 	ServerCheckError(err, "songAlpha: MongoDB connection has failed")
// 	coll := client.Database("songalpha").Collection(alpha)
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	ServerCheckError(err, "songAlpha: allIdx has failed")
// 	var allItems []map[string]string
// 	if err = cur.All(context.TODO(), &allItems); err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allItems)
// }



func main() {
	
=======
func albumInfoByPageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	filter := bson.M{"albumpage": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "initAlbumInfo find has failed")
	var allalbums []AlbVieW2
	if err = cur.All(context.TODO(), &allalbums); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is allalbums", allalbums)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allalbums)
	fmt.Println("Init albumsInfo Complete")
}

func songInfoByPageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	filter := bson.M{"titlepage": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1, "httpaddr": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "ArtPipeline find has failed")
	var tv []map[string]string
	if err = cur.All(context.TODO(), &tv); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is tv", tv)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tv)
	fmt.Println("Initial Song Info Complete")
}

func playSongHandler(w http.ResponseWriter, r *http.Request) {
	fileid := r.URL.Query().Get("selected")
	filter := bson.M{"fileID": fileid}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("maindb").Collection("maindb")
	var results map[string]string
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	var newresults = map[string]string{
		"httpaddr":    results["httpaddr"],
		"artist":      results["artist"],
		"album":       results["album"],
		"title":       results["title"],
		"duration":    results["duration"],
		"picHttpAddr": results["picHttpAddr"],
	}
	fmt.Printf("%s this is playSong", newresults)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newresults)
	fmt.Println("playSong Song Info Complete")
}

func playPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	playlistid := r.URL.Query().Get("selected")
	filter := bson.M{"playlistID": playlistid}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("playlistdb").Collection("playlistdb")
	var results map[string]string
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is playPlaylist", results)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&results)
	fmt.Println("playPlaylist Info Complete")
}

func randomPicsHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "index": 1})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("coverart").Collection("coverart")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "randomPicsHandler has failed")
	var indexliststring []map[string]string
	if err = cur.All(context.TODO(), &indexliststring); err != nil {
		fmt.Println(err)
	}
	var num_list []int
	for _, idx := range indexliststring {
		indexx := idx["index"]
		index1, _ := strconv.Atoi(indexx)
		num_list = append(num_list, index1)
	}
	shuffle(num_list)
	fmt.Println(num_list)
	var randpics []string
	for _, f := range num_list[:12] {
		fmt.Printf("f type: %T", f)
		ff := strconv.Itoa(f)
		fmt.Printf("ff type %T", ff)
		fmt.Printf("ff type %s", ff)
		filter := bson.M{"index": ff}
		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
		defer ampgosetup.Close(client, ctx, cancel)
		ampgosetup.CheckError(err, "MongoDB connection has failed")
		collection := client.Database("coverart").Collection("coverart")
		var rpics map[string]string
		err = collection.FindOne(context.Background(), filter).Decode(&rpics)
		if err != nil {
			fmt.Println(err)
		}
		randpics = append(randpics, rpics["imagehttpaddr"])
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(randpics)
}

///////////////////////////////////////////////////////////////////////////////
//Playlist Stuff
///////////////////////////////////////////////////////////////////////////////
func AllPlayListsFunc() []AmpgoRandomPlaylistData {
	filter := bson.M{}
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
		fmt.Println(err)
	}
	return allplaylists
}

func allPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	allplaylists := AllPlayListsFunc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allplaylists)
}

func deletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	plid := r.URL.Query().Get("playlistid")
	filter := bson.M{"playlistID": plid}
	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
	ServerCheckError(err3, "Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := DeleteOne(client, ctx, "randplaylists", "randplaylists", filter)
	ServerCheckError(err2, "deleteplaylist has failed")
	allplaylists := AllPlayListsFunc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allplaylists)
}

func addPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	plname := r.URL.Query().Get("name")
	plID, _ := UUID()
	var emptymap []map[string]string
	var emptyitem map[string]string = map[string]string{"title": "None Found"}
	emptymap = append(emptymap, emptyitem)
	var plzz AmpgoRandomPlaylistData
	plzz.PlayListName = plname
	plzz.PlayListID = plID
	plzz.PlayList = emptymap
	fmt.Println("This is plzz")
	fmt.Println(plzz)
	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
	ServerCheckError(err3, "Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", &plzz)
	ServerCheckError(err2, "plz insertion has failed")
	allplaylists := AllPlayListsFunc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allplaylists)
}

func shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

// test with curl http://192.168.0.91:9090/CreateRandomPlaylist?songcount=25&&name=RucRandom
func genrandom(maxx int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxx)
}

func addRandomPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	plc := r.URL.Query().Get("songcount")
	plname := r.URL.Query().Get("name")
	fmt.Printf("planame: %s", plname)
	fmt.Printf("plc: %s", plc)
	plcount, _ := strconv.Atoi(plc)
	plID, _ := UUID()
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	coll := client.Database("songtotal").Collection("total")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "allIdx has failed")
	var num_map []map[string]string
	if err = cur.All(context.TODO(), &num_map); err != nil {
		fmt.Println("randplaylist dbcall has fucked up")
		fmt.Println(err)
	}
	fmt.Println(num_map)
	somenum := ""
	for _, item := range num_map {
		somenum = item["total"]
	}
	fmt.Println(somenum)
	var num_list []int
	fmt.Println(plcount)
	for i := 1; i <= plcount; i++ {
		newTotal, _ := strconv.Atoi(somenum)
		fmt.Println(newTotal)
		ranN := genrandom(newTotal)
		fmt.Println(ranN)
		num_list = append(num_list, ranN)
	}
	fmt.Println(num_list)
	shuffle(num_list)
	var randsongs []map[string]string
	for _, f := range num_list {
		if len(randsongs) == plcount {
			break
		}
		ff := strconv.Itoa(f)
		fmt.Println(ff)
		fmt.Printf("this is ff type: %T", ff)
		fmt.Printf("this is f type: %T", f)
		filter := bson.M{"idx": ff}
		client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
		defer ampgosetup.Close(client, ctx, cancel)
		ampgosetup.CheckError(err, "MongoDB connection has failed")
		collection := client.Database("maindb").Collection("maindb")
		var rplaylists map[string]string
		err = collection.FindOne(context.Background(), filter).Decode(&rplaylists)
		if err != nil {
			fmt.Println(ff)
			continue
		} //fmt.Println(err) }
		randsongs = append(randsongs, rplaylists)
	}
	var plz AmpgoRandomPlaylistData
	plz.PlayListName = plname
	plz.PlayListID = plID
	plz.PlayListCount = strconv.Itoa(len(randsongs))
	plz.PlayList = randsongs
	fmt.Println("This is plz")
	fmt.Println(plz)
	client, ctx, cancel, err3 := Connect("mongodb://db:27017/ampgodb")
	ServerCheckError(err3, "Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", &plz)
	ServerCheckError(err2, "plz insertion has failed")
	allplaylists := AllPlayListsFunc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allplaylists)
}

func getCurrentPlayListNameHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{"record": "1"}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "updateCurrentPlayListName: MongoDB connection has failed")
	collection := client.Database("curplaylistname").Collection("curplaylistname")
	var results map[string]string
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&results)
}

func updateCurrentPlayListNameHandler(w http.ResponseWriter, r *http.Request) {
	curplaylistname := r.URL.Query().Get("curplaylistname")
	curplaylistid := r.URL.Query().Get("curplaylistid")
	filter := bson.M{"record": "1"}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "updateCurrentPlayListName: MongoDB connection has failed")
	collection := client.Database("curplaylistname").Collection("curplaylistname")
	var results map[string]string
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	update := bson.M{"$set": bson.M{"curplaylistname": curplaylistname, "curplaylistID": curplaylistid}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	var find_results map[string]string
	err = collection.FindOne(context.Background(), filter).Decode(&find_results)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&find_results)
}

func playListByIDHandler(w http.ResponseWriter, r *http.Request) {
	playlistid := r.URL.Query().Get("playlistid")
	filter := bson.M{"playlistID": playlistid}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "updateCurrentPlayListName: MongoDB connection has failed")
	collection := client.Database("randplaylists").Collection("randplaylists")
	var results AmpgoRandomPlaylistData
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&results)
}

func songInfoFindOne(db string, coll string, filtertype string, filterstring string) map[string]string {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results map[string]string
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}

func playlistInfoFromPlaylistID(db string, coll string, filtertype string, filterstring string) AmpgoRandomPlaylistData {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results AmpgoRandomPlaylistData
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}

func increasePlayListCount(astring string) string {
	newint, _ := strconv.Atoi(astring)
	newnewint := newint + 1
	return strconv.Itoa(newnewint)
}

func decreasePlayListCount(astring string) string {
	newint, _ := strconv.Atoi(astring)
	newnewint := newint - 1
	return strconv.Itoa(newnewint)
}

func addSongToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("fileid")
	plid := r.URL.Query().Get("playlistid")
	fmt.Printf("fileID: %s", fileID)
	fmt.Printf("plid: %s", plid)

	songinfo := songInfoFindOne("maindb", "maindb", "fileID", fileID)
	fmt.Println("This is songinfo")
	fmt.Println(songinfo)

	playlistInfo := playlistInfoFromPlaylistID("randplaylists", "randplaylists", "playlistID", plid)
	fmt.Println("this is playlistinfo")
	fmt.Println(playlistInfo)

	playlistInfo.PlayList = append(playlistInfo.PlayList, songinfo)
	newcount := increasePlayListCount(playlistInfo.PlayListCount)

	update := bson.M{"$set": bson.M{
		"playlistcount": newcount,
		"playlist":      playlistInfo.PlayList,
	},
	}

	filter := bson.M{"playlistID": plid}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("randplaylists").Collection("randplaylists")
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Playlist updated")
}

func deleteSongFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("fileid")
	plid := r.URL.Query().Get("playlistid")
	fmt.Printf("fileID: %s", fileID)
	fmt.Printf("plid: %s", plid)

	PlayListInfo := playlistInfoFromPlaylistID("randplaylists", "randplaylists", "playlistID", plid)
	var newMap []map[string]string
	for _, song := range PlayListInfo.PlayList {
		if song["fileID"] == fileID {
			continue
		} else {
			newMap = append(newMap, song)
		}
	}

	newcount := decreasePlayListCount(PlayListInfo.PlayListCount)

	fmt.Println(newMap)
	fmt.Println(newcount)

	update := bson.M{"$set": bson.M{
		"playlistcount": newcount,
		"playlist":      newMap,
	},
	}

	filter := bson.M{"playlistID": plid}
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "MongoDB connection has failed")
	collection := client.Database("randplaylists").Collection("randplaylists")
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	allplaylists := AllPlayListsFunc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allplaylists)

}

///////////////////////////////////////////////////////////////////////////////
// alphabet stuff
///////////////////////////////////////////////////////////////////////////////

func Unique(arr []string) []string {
	occured := map[string]bool{}
	result := []string{}
	for e := range arr {
		if !occured[arr[e]] {
			occured[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func getMainDbMeta() []map[string]string {
	filter := bson.M{}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "artistFirstLetterHandler: MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "artistFirstLetterHandler: allIdx has failed")
	var letters []map[string]string
	if err = cur.All(context.TODO(), &letters); err != nil {
		fmt.Println(err)
	}
	return letters
}

func artistFirstLetterHandler(w http.ResponseWriter, r *http.Request) {
	theletters := getMainDbMeta()
	var newlist []string
	for _, fl := range theletters {
		fmt.Println(fl)
		boo := fl["artstart"]
		newlist = append(newlist, boo)
	}
	uniquelist := Unique(newlist)
	sort.Strings(uniquelist)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uniquelist)

}

func albumFirstLetterHandler(w http.ResponseWriter, r *http.Request) {
	theletters := getMainDbMeta()
	var newlist []string
	for _, fl := range theletters {
		fmt.Println(fl)
		boo := fl["albstart"]
		newlist = append(newlist, boo)
	}
	uniquelist := Unique(newlist)
	sort.Strings(uniquelist)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uniquelist)

}

func songFirstLetterHandler(w http.ResponseWriter, r *http.Request) {
	theletters := getMainDbMeta()
	var newlist []string
	for _, fl := range theletters {
		fmt.Println(fl)
		boo := fl["titstart"]
		newlist = append(newlist, boo)
	}
	uniquelist := Unique(newlist)
	sort.Strings(uniquelist)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uniquelist)
}

func artistAlphaHandler(w http.ResponseWriter, r *http.Request) {
	alpha := r.URL.Query().Get("alpha")
	filter := bson.M{}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "artistAlpha: MongoDB connection has failed")
	coll := client.Database("artistalpha").Collection(alpha)
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "artistAlpha: allIdx has failed")
	var allItems []map[string]string
	if err = cur.All(context.TODO(), &allItems); err != nil {
		fmt.Println(err)
	}

	if len(allItems) < 1 {
		z := []map[string]string{}
		x := map[string]string{}
		z = append(z, x)
		var noresult ArtVIEW = ArtVIEW{
			Artist:   "None Found",
			ArtistID: "None Found",
			Albums:   z,
			Page:     "None Found",
			Idx:      "None Found",
		}
		zoo := []ArtVIEW{}
		zoo = append(zoo, noresult)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&zoo)
	} else {
		var NewAllItems []ArtVIEW
		for _, item := range allItems {
			filter := bson.M{"artist": item["artist"]}
			client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
			defer ampgosetup.Close(client, ctx, cancel)
			ServerCheckError(err, "artistAlpha: MongoDB connection has failed")
			collection := client.Database("artistview").Collection("artistview")
			var results ArtVIEW
			err = collection.FindOne(context.Background(), filter).Decode(&results)
			if err != nil {
				fmt.Println(err)
			}
			NewAllItems = append(NewAllItems, results)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&NewAllItems)
	}

}

func albumAlphaHandler(w http.ResponseWriter, r *http.Request) {
	alpha := r.URL.Query().Get("alpha")
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "albumAlpha: MongoDB connection has failed")
	coll := client.Database("albumalpha").Collection(alpha)
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "albumAlpha: allIdx has failed")
	var allItems []map[string]string
	if err = cur.All(context.TODO(), &allItems); err != nil {
		fmt.Println(err)
	}

	if len(allItems) < 1 {
		z := []map[string]string{}
		x := map[string]string{}
		z = append(z, x)
		var noresult AlbVieW2 = AlbVieW2{
			Artist:      "None Found",
			ArtistID:    "None Found",
			Album:       "None Found",
			AlbumID:     "None Found",
			Songs:       z,
			AlbumPage:   "None Found",
			NumSongs:    "None Found",
			PicPath:     "None Found",
			Idx:         "None Found",
			PicHttpAddr: "None Found",
		}
		zoo := []AlbVieW2{}
		zoo = append(zoo, noresult)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&zoo)
	} else {
		var NewAllItems []AlbVieW2
		for _, item := range allItems {
			filter := bson.M{"album": item["album"]}
			client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
			defer ampgosetup.Close(client, ctx, cancel)
			ServerCheckError(err, "MongoDB connection has failed")
			collection := client.Database("albumview").Collection("albumview")
			var results AlbVieW2
			err = collection.FindOne(context.Background(), filter).Decode(&results)
			if err != nil {
				fmt.Println(err)
			}
			NewAllItems = append(NewAllItems, results)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&NewAllItems)
	}
}

func songAlphaHandler(w http.ResponseWriter, r *http.Request) {
	alpha := r.URL.Query().Get("alpha")
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := ampgosetup.Connect("mongodb://db:27017/ampgodb")
	defer ampgosetup.Close(client, ctx, cancel)
	ServerCheckError(err, "songAlpha: MongoDB connection has failed")
	coll := client.Database("songalpha").Collection(alpha)
	cur, err := coll.Find(context.TODO(), filter, opts)
	ServerCheckError(err, "songAlpha: allIdx has failed")
	var allItems []map[string]string
	if err = cur.All(context.TODO(), &allItems); err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allItems)
}

func init() {
	ampgosetup.SetUpCheck()
	// var logging_status string = StartServerLogging()
	// fmt.Println(logging_status)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/SetUp", setUpHandler)
	r.HandleFunc("/Home", homeHandler)
	r.HandleFunc("/AlbumsForArtist", albumsForArtistHandler)
	r.HandleFunc("/AlbumsForArtist2", albumsForArtist2Handler)
	r.HandleFunc("/SongsForAlbum", songsForAlbumHandler)
	r.HandleFunc("/RandomPics", randomPicsHandler)

	///////////////////////////////////////////////////////////////////////////

	r.HandleFunc("/AddPlaylist", addPlaylistHandler)
	r.HandleFunc("/AddRandomPlaylist", addRandomPlaylistHandler)
	r.HandleFunc("/AllPlaylists", allPlaylistsHandler)
	r.HandleFunc("/DeletePlayList", deletePlaylistHandler)
	r.HandleFunc("/AddSongToPlaylist", addSongToPlaylistHandler)
	r.HandleFunc("/ArtistInfoByPage", artistInfoByPageHandler)
	r.HandleFunc("/AlbumInfoByPage", albumInfoByPageHandler)
	r.HandleFunc("/SongInfoByPage", songInfoByPageHandler)

	r.HandleFunc("/DeleteSongFromPlaylist", deleteSongFromPlaylistHandler)
	r.HandleFunc("/UpdateCurrentPlayListName", updateCurrentPlayListNameHandler)
	r.HandleFunc("/GetCurrentPlayListName", getCurrentPlayListNameHandler)

	r.HandleFunc("/PlayListByID", playListByIDHandler)

	///////////////////////////////////////////////////////////////////////////
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68

	r := mux.NewRouter()
	// r.HandleFunc("/SetUp", setUpHandler)
	r.HandleFunc("/Home", homeHandler)
	// r.HandleFunc("/AlbumsForArtist", albumsForArtistHandler)
	// r.HandleFunc("/AlbumsForArtist2", albumsForArtist2Handler)
	// r.HandleFunc("/SongsForAlbum", songsForAlbumHandler)
	// r.HandleFunc("/RandomPics", randomPicsHandler)

	///////////////////////////////////////////////////////////////////////////

	// r.HandleFunc("/AddPlaylist", addPlaylistHandler)
	// r.HandleFunc("/AddRandomPlaylist", addRandomPlaylistHandler)
	// r.HandleFunc("/AllPlaylists", allPlaylistsHandler)
	// r.HandleFunc("/DeletePlayList", deletePlaylistHandler)
	// r.HandleFunc("/AddSongToPlaylist", addSongToPlaylistHandler)
	// r.HandleFunc("/ArtistInfoByPage", artistInfoByPageHandler)
	// r.HandleFunc("/AlbumInfoByPage", albumInfoByPageHandler)
	// r.HandleFunc("/SongInfoByPage", songInfoByPageHandler)

	// r.HandleFunc("/DeleteSongFromPlaylist", deleteSongFromPlaylistHandler)
	// r.HandleFunc("/UpdateCurrentPlayListName", updateCurrentPlayListNameHandler)
	// r.HandleFunc("/GetCurrentPlayListName", getCurrentPlayListNameHandler)

	// r.HandleFunc("/PlayListByID", playListByIDHandler)

	// ///////////////////////////////////////////////////////////////////////////

	// r.HandleFunc("/PlaySong", playSongHandler)
	// r.HandleFunc("/PlayPlaylist", playPlaylistHandler)

	// ///////////////////////////////////////////////////////////////////////////

	// r.HandleFunc("/ArtistAlpha", artistAlphaHandler)
	// r.HandleFunc("/AlbumAlpha", albumAlphaHandler)
	// r.HandleFunc("/SongAlpha", songAlphaHandler)

	r.HandleFunc("/ArtistFirstLetter", artistFirstLetterHandler)
	r.HandleFunc("/AlbumFirstLetter", albumFirstLetterHandler)
	r.HandleFunc("/SongFirstLetter", songFirstLetterHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/root/static/"))))
	r.PathPrefix("/fsData/").Handler(http.StripPrefix("/fsData/", http.FileServer(http.Dir("/root/fsData/"))))
<<<<<<< HEAD
	

	http.ListenAndServe(":9090",  handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
=======
	http.ListenAndServe(":9090", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r))


	// http.ListenAndServeTLS(":443", "cert.pem", "key.pem", 
	// 	handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	// 	handlers.AllowedOrigins([]string{"*"}))(r))
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func DeleteOne(client *mongo.Client, ctxx context.Context, dataBase, col string, doc interface{}) (*mongo.DeleteResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result2, err1 := collection.DeleteOne(ctxx, doc)
	return result2, err1
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
