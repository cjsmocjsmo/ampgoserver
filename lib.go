package main

import (
	// "context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"

	// "github.com/bogem/id3v2"
	"github.com/disintegration/imaging"
	// "go.mongodb.org/mongo-driver/bson"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "io/ioutil"
	"log"
	"os"

	// "path/filepath"
	"strconv"

	"strings"
	// "time"
)

func ConvertSTR(astring string) int {
	Ofset, err := strconv.Atoi(astring)
	CheckError(err, "strconv has failed")
	return Ofset
}

func CheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err)
		log.Println(msg)
		log.Println(err)
		panic(err)
	}
}

func StartSetupLogging() string {
	logtxtfile := os.Getenv("AMPGO_SETUP_LOG_PATH")
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logtxtfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file)
	fmt.Println("Logging started")
	return "server logging started"
}

func Read_File_mp3(apath string) {
	var jsonmp3 JsonMP3
	data, er := os.ReadFile(apath)
	CheckError(er, "ReadFile has fucked up")
	err := json.Unmarshal(data, &jsonmp3)
	CheckError(err, "unmarshal has fucked up")
	// fmt.Println(jsonmp3)
	InsertMP3Json("maindb", "mp3s", jsonmp3)
	// log.Printf("%s : Read_File_mp3 complete", apath)
}

func Read_File_jpg(apath string) {
	var jsonjpg JsonJPG
	data, er := os.ReadFile(apath)
	CheckError(er, "jpg readfile has fucked up")
	err := json.Unmarshal(data, &jsonjpg)
	CheckError(err, "jpg unmarshal has fucked up")
	// fmt.Println(jsonjpg)
	InsertJPGJson("maindb", "jpgs", jsonjpg)
	// log.Printf("%s : Read_File_jpg complete", apath)
}

func Read_File_pages(apath string) {
	var jsonpages JsonPage
	data, er := os.ReadFile(apath)
	CheckError(er, "file pages has fucked up")
	err := json.Unmarshal(data, &jsonpages)
	CheckError(err, "file pages unmarshal has fucked up")
	// fmt.Println(jsonpages)
	InsertPagesJson("maindb", "pages", jsonpages)
	// log.Printf("%s : Read_File_pages complete", apath)
}

func Read_Artist_ID_List(apath string) {
	var artids ArtID
	data, er := os.ReadFile(apath)
	CheckError(er, "Read_Artist_ID_List has fucked up")
	err := json.Unmarshal(data, &artids)
	CheckError(err, "Read_Artist_ID_List unmarshal has fucked up")
	// fmt.Println(jsonpages)
	for _, aid := range artids.ArtistIDList {
		voodoo := make(map[string]string)
		voodoo["Artist"] = aid.Artist
		voodoo["ArtistID"] = aid.ArtistID
		AmpgoInsertOne("ids", "artistids", voodoo)
	}
	// log.Printf("%s : Read_File_pages complete", apath)
}

func Read_Album_ID_List(apath string) {
	var albids AlbID
	data, er := os.ReadFile(apath)
	CheckError(er, "Read_Album_ID_List has fucked up")
	err := json.Unmarshal(data, &albids)
	CheckError(err, "Read_Album_ID_List unmarshal has fucked up")
	// fmt.Println(jsonpages)
	for _, aaid := range albids.AlbumIDList {
		vd := make(map[string]string)
		vd["Album"] = aaid.Album
		vd["AlbumID"] = aaid.AlbumID
		AmpgoInsertOne("ids", "albumids", vd)
	}
	InsertAlbumIDSJson("ids", "albumids", albids)
	// log.Printf("%s : Read_File_pages complete", apath)
}



func UpdateMainDB(m2 JsonMP3) {
	artid := gArtistInfo(m2.Tags_artist)
	artID := artid["ArtistID"]
	fmt.Println("this is artID")
	fmt.Println(artID)

	albid := gAlbumInfo(m2.Tags_album)
	albID := albid["AlbumID"]
	fmt.Println("this is albID")
	fmt.Println(albID)

	Doko := make(map[string]string)
	Doko["Dir"] = m2.Dir
	Doko["Full_Filename"] = m2.Full_Filename
	Doko["Filename"] = m2.Filename
	Doko["Ext"] = m2.Ext
	Doko["File_id"] = m2.File_id
	Doko["File_Size"] = m2.File_Size
	Doko["Artist"] = m2.Tags_artist
	Doko["ArtistID"] = artID
	Doko["Artist_first"] = strings.ToUpper(m2.Artist_first)
	Doko["Album"] = m2.Tags_album
	Doko["AlbumID"] = albID
	Doko["Album_first"] = strings.ToUpper(m2.Album_first)
	Doko["Song"] = m2.Tags_song
	Doko["SongID"] = m2.File_id
	Doko["Song_first"] = strings.ToUpper(m2.Song_first)
	// Doko.Genre = m2["genre"]
	Doko["Index"] = m2.Index
	Doko["Play_length"] = m2.Play_length
	Doko["Img_base64_str"] = m2.Img_base64_str
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "UpdateMainDB: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "maindb", "maindb", &Doko)
	CheckError(err2, "UpdateMainDB: maindb insertion has failed")
	return
}

func GDistArtistForAgg() (dArtAll []map[string]string) {
	dArtist := AmpgoDistinct("maindb", "maindb", "Artist")
	for _, art := range dArtist {
		dArt := AmpgoFindOne("maindb", "maindb", "Artist", art)
		dArtAll = append(dArtAll, dArt)
	}
	return dArtAll
}

func create_just_albumID_list(alist []map[string]string) (just_albumID_list []string) {
	for _, albID := range alist {
		just_albumID_list = append(just_albumID_list, albID["AlbumID"])
	}
	return
}

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

func get_albums_for_artist(fullalblist []map[string]string) (final_alblist []map[string]string) {
	just_albumID_list := create_just_albumID_list(fullalblist)
	//remove double albumid entries
	unique_items := Unique(just_albumID_list)
	for _, uitem := range unique_items {
		albINFO := AmpgoFindOne("maindb", "maindb", "AlbumID", uitem)
		final_alblist = append(final_alblist, albINFO)
	}
	return
}

func ArtPipline(artmap map[string]string, page int, idx int) (MyArView ArtVieW2) {
	dirtyalblist := AmpgoFind("maindb", "maindb", "ArtistID", artmap["ArtistID"]) //[]map[string]string
	results2 := get_albums_for_artist(dirtyalblist)
	albc := len(results2)
	albcount := strconv.Itoa(albc)
	MyArView.Artist = artmap["Artist"]
	MyArView.ArtistID = artmap["ArtistID"]
	MyArView.Albums = results2
	MyArView.AlbCount = albcount
	MyArView.Page = strconv.Itoa(page)
	MyArView.Index = strconv.Itoa(idx)
	return
}

func GDistAlbum() (DAlbAll []map[string]string) {
	DAlbumID := AmpgoDistinct("maindb", "maindb", "AlbumID")
	for _, albID := range DAlbumID {
		DAlb := AmpgoFindOne("maindb", "maindb", "AlbumID", albID)
		DAlbAll = append(DAlbAll, DAlb)
	}
	return
}

func get_songs_for_album(fullsonglist []map[string]string) (final_songlist []map[string]string) {
	//a list of just albumid's
	var just_songID_list []string
	for _, song := range fullsonglist {
		just_songID_list = append(just_songID_list, song["File_id"])
	}
	//remove double songID entries
	unique_items := Unique(just_songID_list)
	for _, uitem := range unique_items {
		songINFO := AmpgoFindOne("maindb", "maindb", "File_id", uitem)
		si := make(map[string]string)
		si["Index"] = songINFO["Index"]
		si["Filename"] = songINFO["Filename"]
		si["ArtistID"] = songINFO["ArtistID"]
		si["AlbumID"] = songINFO["AlbumID"]
		si["SongID"] = songINFO["SongID"]
		si["Dir"] = songINFO["Dir"]
		si["Ext"] = songINFO["Ext"]
		si["Play_length"] = songINFO["Play_length"]
		si["Full_Filename"] = songINFO["Full_Filename"]
		si["File_id"] = songINFO["File_id"]
		si["File_Size"] = songINFO["File_Size"]
		si["Album_first"] = songINFO["Album_first"]
		si["Song_first"] = songINFO["Song_first"]
		si["Artist"] = songINFO["Artist"]
		si["Artist_first"] = songINFO["Artist_first"]
		si["Album"] = songINFO["Album"]
		si["Song"] = songINFO["Song"]
		final_songlist = append(final_songlist, si)
	}
	return final_songlist
}

func AlbPipeline(DAlb map[string]string, page int, idx int) (MyAlbview AlbVieW2) {
	dirtysonglist := AmpgoFind("maindb", "maindb", "AlbumID", DAlb["AlbumID"])
	results := get_songs_for_album(dirtysonglist)
	songcount := len(results)
	MyAlbview.Artist = DAlb["Artist"]
	MyAlbview.ArtistID = DAlb["ArtistID"]
	MyAlbview.Album = DAlb["Album"]
	MyAlbview.AlbumID = DAlb["AlbumID"]
	MyAlbview.NumSongs = strconv.Itoa(songcount)
	MyAlbview.Img_base64_str = DAlb["Img_base64_str"]
	MyAlbview.Songs = results
	MyAlbview.AlbumPage = strconv.Itoa(page)
	MyAlbview.Idx = strconv.Itoa(idx)
	// MyAlbview.PicHttpAddr = DAlb["picHttpAddr"]
	return
}

//////////////////////////////////////////////////////////////////////////

func getFileInfo(apath string) (filename string, size string) {
	ltn, err := os.Open(apath)
	CheckError(err, "getFileInfo: file open has fucked up")
	defer ltn.Close()
	ltnInfo, _ := ltn.Stat()
	filename = ltnInfo.Name()
	size = strconv.FormatInt(ltnInfo.Size(), 10)
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

func resizeImage(infile string, outfile string) string {
	pic, err := imaging.Open(infile)
	if err != nil {
		return os.Getenv("AMPGO_NO_ART_PIC_PATH")
	}
	sjImage := imaging.Resize(pic, 200, 0, imaging.Lanczos)
	err = imaging.Save(sjImage, outfile)
	CheckError(err, "resizeImage: image save has fucked up")
	return outfile
}

/////////////////////////////////////////////////////////////////////////////////////////////

// func CreateRandomPlaylistDB() string {
// 	var ranDBInfo randDb
// 	var emptylist []map[string]string
// 	var emptyitem map[string]string = map[string]string{"None": "No Songs Found"}
// 	emptylist = append(emptylist, emptyitem)
// 	uuid, _ := UUID()
// 	ranDBInfo.PlayListName = "EmptyRandomPlaylist"
// 	ranDBInfo.PlayListID = uuid
// 	ranDBInfo.PlayListCount = "0"
// 	ranDBInfo.Playlist = emptylist
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
// 	CheckError(err, "CreateRandomPlaylistDB: Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, "randplaylists", "randplaylists", ranDBInfo)
// 	CheckError(err2, "CreateRandomPlaylistDB: randplaylists insertion has failed")
// 	return "Created"
// }

// func CreateCurrentPlayListNameDB() string {
// 	var curPlayListName map[string]string = map[string]string{"record": "1", "curplaylistname": "None", "curplaylistID": "None"}
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
// 	CheckError(err, "InsertDurationInfo: Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, "curplaylistname", "curplaylistname", &curPlayListName)
// 	CheckError(err2, "InsertDurationInfo: curplaylistname insertion has failed")
// 	return "curplaylistname Created"
// }
