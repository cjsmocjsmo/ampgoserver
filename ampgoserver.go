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
	// "fmt"
	"log"
	// "math/rand"
	"os"
	// "path"
	// "strconv"
	// "strings"
	// "io/ioutil"
	"net/http"
	"html/template"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/handlers"
	// ampgosetup "github.com/cjsmocjsmo/ampgosetup"
	// "github.com/cjsmocjsmo/ampgosetup"
	"github.com/cjsmocjsmo/ampgosetup"
	// "ampgosetup"
)

const (
	OffSet = 10
)

type plist struct {
	PLName string              `bson:"PLName"`
	PLId   string              `bson:"PLId"`
	Songs  []map[string]string `bson:"Songs"`
}

type iMgfa struct {
	Album   string
	HSImage string
	Songs   []map[string]string
}

type rAlbinfo struct {
	Songs   []map[string]string `bson:"songs"`
	HSImage string
}

type voodoo struct {
	Playlists []map[string]string
}

func sfdbCon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("AMP_AMPDB_ADDR"))
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}


///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

//ArtVIEW exported
type ArtVIEW struct {
	Artist   string              `bson:"artist"`
	ArtistID string              `bson:"artistID"`
	Albums   []map[string]string `bson:"albums"`
	Page     int                 `bson:"page"`
	Idx      int                 `bson:"idx"`
}

func initialArtistInfoHandler(w http.ResponseWriter, r *http.Request) {
	ofset := OffSet
	ses := sfdbCon()
	defer ses.Close()
	AMPc := ses.DB("artistview").C("artistviews")
	b1 := bson.M{"_id": 0}
	var av []ArtVIEW
	err := AMPc.Find(nil).Select(b1).Sort("artist").Limit(ofset).All(&av)
	if err != nil {
		log.Println("find one has failed")
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&av)
	log.Println("Initial Artist Info Complete")
}

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

func initialalbumInfoHandler(w http.ResponseWriter, r *http.Request) {
	ofset := OffSet
	ses := sfdbCon()
	defer ses.Close()
	ALBc := ses.DB("albview").C("albview")
	b1 := bson.M{"_id": 0}
	var albv []AlbvieW
	err := ALBc.Find(nil).Select(b1).Sort("album").Limit(ofset).All(&albv)
	if err != nil {
		log.Println("initial album info has fucked up")
		log.Println(err)
	}
	log.Println("GInitialAlbumInfo is complete")
	log.Println(&albv)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&albv)
	log.Println("Initial Artist Info Complete")
}

func initialsongInfoHandler(w http.ResponseWriter, r *http.Request) {
	ofset := OffSet
	ses := sfdbCon()
	defer ses.Close()
	MAINc := ses.DB("maindb").C("maindb")
	b1 := bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1}
	var tv []map[string]string
	err := MAINc.Find(nil).Select(b1).Limit(ofset).Sort("title").All(&tv)
	if err != nil {
		log.Println("intial song info fucked up")
		log.Println(err)
	}
	log.Println(&tv)
	log.Println("GInitialSongInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tv)
	log.Println("Initial Artist Info Complete")
}

func artistPageHandler(w http.ResponseWriter, r *http.Request) {
	ses := sfdbCon()
	defer ses.Close()
	ARTVc := ses.DB("artistview").C("artistviews")
	// var ARDist []map[string]string
	var ARDist []string
	err := ARTVc.Find(nil).Distinct("page", &ARDist)
	if err != nil {
		log.Println("artist alpha has fucked up")
		log.Println(err)
	}
	log.Println("ArtistAlpha is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ARDist)
}

func albumPageHandler(w http.ResponseWriter, r *http.Request) {
	ses := sfdbCon()
	defer ses.Close()
	ALBVc := ses.DB("albview").C("albview")
	// var ALDist []AlbvieW
	var ALDist []string
	err := ALBVc.Find(nil).Distinct("page", &ALDist)
	if err != nil {
		log.Println("album alpha fucked up")
		log.Println(err)
	}
	log.Println("AlbumAlpha is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ALDist)
}

func titlePageHandler(w http.ResponseWriter, r *http.Request) {
	ses := sfdbCon()
	defer ses.Close()
	MAINc := ses.DB("maindb").C("maindb")
	// var TDist []map[string]string
	var TDist []string
	err := MAINc.Find(nil).Distinct("page", &TDist)
	if err != nil {
		log.Println("title alpha fucked up")
		log.Println(err)
	}
	log.Println("TitleAlpha is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&TDist)
}

func imageSongsForAlbumHandler(w http.ResponseWriter, r *http.Request) {
	albid := r.URL.Query().Get("selected")
	ses := sfdbCon()
	defer ses.Close()
	ALBc := ses.DB("albview").C("albview")
	b1 := bson.M{"albumID": albid}
	b2 := bson.M{"_id": 0, "album": 1, "songs": 1, "hsimage": 1}
	var iM []iMgfa
	err := ALBc.Find(b1).Select(b2).One(&iM)
	if err != nil {
		log.Println("gimage song for album fucked up")
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&iM)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	// ST := ampgolib.GStats()
	ses := sfdbCon()
	defer ses.Close()
	STATc := ses.DB("goampgo").C("dbstats")
	b1 := bson.M{"_id": 0}
	var st map[string]string
	err := STATc.Find(nil).Select(b1).One(&st)
	if err != nil {
		log.Println("stats has fucked up")
		log.Println(err)
	}
	log.Println("GStats is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&st)
}

// func randomPicsHandler(w http.ResponseWriter, r *http.Request) {
// 	RandomPics := ampgolib.RPics()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(RandomPics)
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

func songInfoHandler(w http.ResponseWriter, r *http.Request) {
	pagenum := r.URL.Query().Get("selected")
	ses := sfdbCon()
	defer ses.Close()
	MAINc := ses.DB("maindb").C("maindb")
	b1 := bson.M{"page": pagenum}
	b2 := bson.M{"_id": 0, "title": 1, "fileID": 1, "artist": 1}
	var SIS []map[string]string
	err := MAINc.Find(b1).Select(b2).All(&SIS)
	if err != nil {
		log.Println("song info has fucked up")
		log.Println(err)
	}
	log.Println("SongInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SIS)
}

func albumInfoHandler(w http.ResponseWriter, r *http.Request) {
	pagenum := r.URL.Query().Get("selected")
	ses := sfdbCon()
	defer ses.Close()
	ALBVc := ses.DB("albview").C("albview")
	b1 := bson.M{"page": pagenum}
	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
	var AI []AlbvieW
	err := ALBVc.Find(b1).Select(b2).All(&AI)
	if err != nil {
		log.Println("AlbumInfo has fucked up")
		log.Println(err)
	}
	log.Println("AlbumInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&AI)
}

func artistInfoHandler(w http.ResponseWriter, r *http.Request) {
	pagenum := r.URL.Query().Get("selected")
	ses := sfdbCon()
	defer ses.Close()
	ARTc := ses.DB("artistview").C("artistviews")
	b1 := bson.M{"page": pagenum}
	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "albums": 1, "page": 1}
	var ARTI []ArtVIEW
	err := ARTc.Find(b1).Select(b2).All(&ARTI)
	if err != nil {
		log.Println("ArtistInfo has fucked up")
		log.Println(err)
	}
	log.Println("ArtistInfo is complete")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&ARTI)
}

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

func setUpHandler(w http.ResponseWriter, r *http.Request) {
	ampgosetup.Setup()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Setup Complete")
}

// func setUpHandler(w http.ResponseWriter, r *http.Request) {
// 	ampgolib.SetUp()
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Setup Complete")
// }

// func introHandler(w http.ResponseWriter, r *http.Request) {
// 	var p = map[string]string{"Title" : "AmpGo"}
//     t := template.Must(template.ParseFiles("assets/templates/intro.html"))
//     t.Execute(w, p)
// }


func homeHandler(w http.ResponseWriter, r *http.Request) {
	var p = map[string]string{"Title" : "AmpGo"}
    t := template.Must(template.ParseFiles("assets/templates/home.html"))
    t.Execute(w, p)
}

func init() {
	ampgosetup.Setup()
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/assets").Subrouter()
	// r.HandleFunc("/Intro", introHandler)
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

	// r.HandleFunc("/RandomPics", randomPicsHandler)
	// r.HandleFunc("/RamdomAlbumPicPlaySong", ramdomAlbumPicPlaySongHandler)
	r.HandleFunc("/Stats", statsHandler)
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
	// r.HandleFunc("/SetUp", setUpHandler)
	
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/static/"))))
	http.ListenAndServe(":9090", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), 
		handlers.AllowedOrigins([]string{"*"}))(r))


}
