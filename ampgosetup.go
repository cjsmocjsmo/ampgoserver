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
	"strings"
	"log"
	"os"
	// "path"
	"runtime"

	"sync"
	"time"
	// "context"
	// "encoding/json"
	"path/filepath"
	// "strconv"
	// "gopkg.in/yaml.v3"
)

// var OFFSET string = os.Getenv("AMPGO_OFFSET")
var OffSet int = ConvertSTR(OFFSET)

func SetUp() {
	StartSetupLogging()
	ti := time.Now()
	fmt.Println(ti)
	log.Println(ti)
	runtime.GOMAXPROCS(runtime.NumCPU())
	var addr string = os.Getenv("AMPGO_MEDIA_METADATA_PATH")
	var address string = addr + "/*.json"
	log.Println(address)
	files, err := filepath.Glob(address)
	if err != nil {
        log.Println(err)
    }

	log.Println("starting walk")
	for _, foo := range files {
		switch{
		case strings.Contains(foo, "mp3"):
			// log.Println(idx, foo)
			Read_File_mp3(foo)

		case strings.Contains(foo, "jpg"):
			// log.Println(idx, foo)
			Read_File_jpg(foo)

		case strings.Contains(foo, "page"):
			// log.Println(idx, foo)
			Read_File_pages(foo)
		}
	}
	log.Println("walk is complete")

	log.Println("starting GetDistAlbumMeta1")
	// dalb := AmpgoDistinct("tempdb1", "meta1", "album")
	dalb := AmpgoDistinct("maindb", "mp3s", "tags_album")
	// log.Println(dalb)
	// fmt.Println(dalb)
	log.Println("GetDistAlbumMeta1 is complete ")

	log.Println("starting InsAlbumID")
	var wg1 sync.WaitGroup
	for _, alb := range dalb {
		wg1.Add(1)
		go func(alb string) {
			InsAlbumID(alb)
			wg1.Done()
		}(alb)
		wg1.Wait()
	}
	log.Println("InsAlbumID is complete ")

	log.Println("starting GDistArtist")
	// dart := AmpgoDistinct("tempdb1", "meta1", "artist")
	dart := AmpgoDistinct("maindb", "mp3s", "tags_artist")
	// log.Println(dart)
	log.Println("GDistArtist is complete ")

	log.Println("starting InsArtistID")
	var wg2 sync.WaitGroup
	for _, art := range dart {
		wg2.Add(1)
		go func(art string) {
			InsArtistID(art)
			wg2.Done()
		}(art)
		wg2.Wait()
	}
	log.Println("InsArtistID is complete ")

	log.Println("starting GetAllObjects")
	AllObjs := GetAllObjects()
	// log.Println(AllObjs)
	log.Println("GetAllObjects is complete ")
	

	log.Println("starting UpdateMainDB")
	var wg3 sync.WaitGroup
	for _, blob := range AllObjs {
		// log.Println(blob)
		wg3.Add(1)
		go func(blob JsonMP3) {
			UpdateMainDB(blob)
			wg3.Done()
		}(blob)
		wg3.Wait()
	}
	log.Println("UpdateMainDB is complete ")

	log.Println("starting ArtistFirst ")
	var wg99a sync.WaitGroup
	for _, art := range dart {
		wg99a.Add(1)
		go func(art string) {
			ArtistFirst(art)
			wg99a.Done()
		}(art)
		wg99a.Wait()
	}
	log.Println("ArtistFirst is complete ")

	log.Println("starting AlbumFirst ")
	var wg99 sync.WaitGroup
	for _, alb := range dalb {
		wg99.Add(1)
		go func(alb string) {
			AlbumFirst(alb)
			wg99.Done()
		}(alb)
		wg99.Wait()
	}
	log.Println("AlbumFirst is complete ")

	SongFirst()

	// // //AggArtist
	log.Println("starting GDistArtistForAgg")
	DistArtist := GDistArtistForAgg()
	// log.Println(DistArtist)
	fmt.Println(DistArtist)
	log.Println("GDistArtistForAgg is complete ")

	fmt.Println("starting AggArtist")
	var wg5 sync.WaitGroup
	var artpage int = 0
	for artIdsx, DArtt := range DistArtist {
		if artIdsx < OffSet {
			artpage = 1
		} else if artIdsx%OffSet == 0 {
			artpage++
		} else {
			artpage = artpage + 0
		}

		APL := ArtPipline(DArtt, artpage, artIdsx)

		wg5.Add(1)
		go func(APL ArtVieW2) {
			InsArtPipeline(APL)
			wg5.Done()
		}(APL)
		wg5.Wait()
	}
	fmt.Println("AggArtists is complete")
	log.Println("AggArtists is complete")
	

	//AggAlbum
	log.Println("AggAlbum has started")

	log.Println("Starting GDistAlbum3")
	DistAlbum := GDistAlbum()

	var wg6 sync.WaitGroup
	var albpage int = 0
	for albIdx, DAlb := range DistAlbum {
		wg6.Add(1)
		if albIdx < OffSet {
			albpage = 1
		} else if albIdx%OffSet == 0 {
			albpage++
		} else {
			albpage = albpage + 0
		}
		APLX := AlbPipeline(DAlb, albpage, albIdx)
		log.Println("this is aplx")
		log.Println(APLX)
		go func(APLX AlbVieW2) {
			InsAlbViewID(APLX)
			wg6.Done()
		}(APLX)
		wg6.Wait()
	}
	
	// CreateRandomPicsDB()

	// CreateRandomPlaylistDB()

	// CreateCurrentPlayListNameDB()


	t2 := time.Now().Sub(ti)
	fmt.Println(t2)
	fmt.Println("THE END")
}
