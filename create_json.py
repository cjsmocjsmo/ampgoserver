#!/usr/bin/env python3

import hashlib
import os
import shutil
import json
import glob
import base64

from PIL import Image
import mutagen
from mutagen import File
from mutagen.mp3 import MP3

from pprint import pprint


class MP3Tags:
	Track = None
	Artist = None
	Album = None
	Song = None
	
	def __init__(self, fn):
		self.fn = fn

		try:
			self.audio = File(self.fn)
		except (KeyError, mutagen.mp3.HeaderNotFoundError, AttributeError):
			print(self.fn)
			pass

		try:
			type(self).Track = self.audio['TRCK'].text[0]
		except (KeyError, AttributeError):
			type(self).Track = '50'
	
		try:
			type(self).Artist = self.audio["TPE1"].text[0]
		except (KeyError, AttributeError): 
			type(self).Artist = 'Fuck Artist'
			print(''.join(("KeyError: No TPE1 tag... ", self.fn)))
	
		try:
			type(self).Album = self.audio["TALB"].text[0]
		except (KeyError, AttributeError): 
			type(self).Album = 'Fuck Album'
			print(''.join(("KeyError No TALB tag ... ", self.fn)))
			
		try:
			type(self).Song = self.audio['TIT2'].text[0]
		except (KeyError, AttributeError): 
			type(self).Song = 'Fuck Song'
			print(''.join(("KeyError: No TIT2 tag... ", self.fn)))





class Artist:
    def __init__(self):
        self.apath = "/media/charliepi/FOO/music/music"  #os.getenv('AMPGO_MEDIA_PATH')
        self.basedir = "/media/charliepi/FOO/music/music"  #os.getenv('AMPGO_MEDIA_PATH')
        
    def artist_check(self, alist):
        count = 0
        for meta in alist:
            if meta['Ext'] == ".mp3":
                da = meta['Dir_artist'].replace("_", " ")
                fa = meta['File_artist'].replace("_", " ")
                ta = meta['Tags_artist']
                if da == fa and fa == ta:
                    pass
                else:
                    count += 1
                    print(da, fa, ta)
                    print(meta['Full_Filename'])
        print(count)

    def album_check(self, alist):
        count = 0
        for meta in alist:
            if meta['Ext'] == ".mp3":
                da = meta['Dir_album'].replace("_", " ")
                fa = meta['File_album'].replace("_", " ")
                ta = meta['Tags_album']
                if da == fa and fa == ta:
                    pass
                else:
                    count += 1
                    print(da, fa, ta)
                    print(meta['Full_Filename'])
        print(count)

    def song_check(self, alist):
        count = 0
        for meta in alist:
            if meta['Ext'] == ".mp3":
                ds = meta['File_song'].replace("_", " ")
                fs = meta['Tags_song'].replace("_", " ")
                
                if ds == fs:
                    pass
                else:
                    count += 1
                    print(ds, fs)
                    print(meta['Full_Filename'])
        print(count)

    def glob_for_pics(self, dir_list):
        missing_pic_list = []
        for dir in dir_list:
            globpath = dir + "/*.jpg"
            globpath2 = dir + "/*.png"
            result = glob.glob(globpath)
            result2 = glob.glob(globpath2)
            if len(result) == 0 and len(result2) == 0:
                missing_pic_list.append(dir)
        return (list(set(missing_pic_list)), len(missing_pic_list))


class MusicFiles:
    def __init__(self):


        self.apath = "/media/charliepi/FOO/music/music"
        self.BaseDir = "/media/charliepi/FOO/music/music"
        self.metaDir = "/media/charliepi/FOO/music/metadata/"
        self.Dirlist = []
        self.Artist_List = []
        self.Album_List = []
        self.Song_List = []

    def clean_metaDir(self):
        shutil.rmtree(self.metaDir)
        os.mkdir(self.metaDir)

    def check_jpg(self, afile):
        Dir, Filename = os.path.split(afile)
        jpgpath = Dir + "/Folder.jpg"
        if os.path.exists(jpgpath):
            return True, jpgpath
        else:
            return False

    def img_size(self, afile):
        with Image.open(afile) as img_s:
            width, height = img_s.size
            return (width, height)

    def img_to_base64(self, afile):
        with open(afile, "rb") as img_file:
            my_string = base64.b64encode(img_file.read())
        return my_string.decode('utf-8')

    def calc_md5(self, afile):
        md5 = hashlib.md5(afile.encode('utf-8')).hexdigest()
        return md5
        
    def play_length(self, afile):
        audio = MP3(afile)
        audio_info = audio.info
        length_in_secs = int(audio_info.length)
        length_in_mills = length_in_secs * 1000
        return length_in_mills

    def metadata_from_file(self, afile, acount):
        meta = {}
        meta['BaseDir'] = self.BaseDir
        meta['Full_Filename'] = afile

        size = os.path.getsize(afile)
        meta['File_Size'] = str(size)

        thefile, ext = os.path.splitext(afile)
        meta["Ext"] = ext

        fullDir, Filename = os.path.split(thefile)
        meta["Filename"] = Filename

        _, Dir = fullDir.split(self.BaseDir)
        meta['Dir'] = Dir

        dsplitlist = Dir.split("/")[1:]
        #meta['Dir_Split_List'] = dsplitlist
        meta['Dir_catagory'] = dsplitlist[0]
        meta['Dir_artist'] = dsplitlist[1]
        meta['Dir_album'] = dsplitlist[2]

        meta['Index'] = str(acount)
        meta['Dir_delem'] = "/"

        fid = str(self.calc_md5(afile))
        meta['File_id'] = fid
        if ext != ".mp3":
            width, height = self.img_size(afile)
            meta['Jpg_width'] = str(width)
            meta['Jpg_height'] = str(height)
            meta['File_delem'] = "None"
            meta["Img_base64_str"] = self.img_to_base64(afile)
            return meta
        else:
            meta['File_delem'] = "_-_"
            File_split_list = Filename.split("_-_")
            #meta['File_split_list'] = File_split_list
            meta['Track'] = File_split_list[0]
            meta['File_artist'] = File_split_list[1]
            meta['File_album'] = File_split_list[2]
            meta['File_song'] = File_split_list[3]
            tags = MP3Tags(afile)
            meta['Tags_artist'] = tags.Artist
            self.Artist_List.append(tags.Artist)

            meta['Tags_album'] = tags.Album
            self.Album_List.append(tags.Album)

            meta['Tags_song'] = tags.Song
            self.Song_List.append(tags.Song)

            meta['Artist_first'] = tags.Artist[:1]
            meta['Album_first'] = tags.Album[:1]
            meta['Song_first'] = tags.Song[:1]
            meta['Play_length'] = str(self.play_length(afile))
            boo = self.check_jpg(afile)
            if boo[0]:
                meta["Img_base64_str"] = self.img_to_base64(boo[1])
            else:
                meta["Img_base64_str"] = "None"
        return meta

    def write_to_file(self, meta, acount):
        if meta['Ext'] == ".mp3":
            newfile = self.metaDir + str(acount) + "_mp3_data.json"
            with open(newfile, "w") as nf:
                data = json.dumps(meta, indent = 4)
                nf.write(data)
            print("\n")
            print(newfile)
            print("\n")
        else:
            newfile = self.metaDir + str(acount) + "_jpg_data.json"
            with open(newfile, "w") as nf:
                data = json.dumps(meta, indent = 4)
                nf.write(data)
            print("\n")
            print(newfile)
            print("\n")

    def find(self):
        all_files_list = []
        count = 0
        for (paths, Dirs, files) in os.walk(self.apath, followlinks=True):
            for Filename in files:
                fnn = os.path.join(paths, Filename)
                Dir, _ = os.path.split(fnn)
                _, ext = os.path.splitext(fnn)
                if ext == ".sh":
                    pass
                else:
                    count += 1
                    self.Dirlist.append(Dir)
                    meta = self.metadata_from_file(fnn, count)
                    
                    self.write_to_file(meta, count)
                    all_files_list.append(meta)
        return all_files_list

    def get_lists(self):
        new_artist_file = "/media/charliepi/FOO/music/metadata/Artist_ID_List.json"
        newArtistList = list(set(self.Artist_List))
        newArtistList.sort()

        ArtistIdList = []
        for artist in newArtistList:
            artistID = str(self.calc_md5(artist))
            foobar = {
                "Artist": artist,
                "ArtistID": artistID,
            }
            ArtistIdList.append(foobar)
        newArtiistDict = {
            "ArtistIDList": ArtistIdList,
        }            
        with open(new_artist_file, "w") as nf:
            data = json.dumps(newArtiistDict, indent = 4)
            nf.write(data)


        new_album_file = "/media/charliepi/FOO/music/metadata/Album_ID_List.json"
        newAlbumList = list(set(self.Album_List))
        newAlbumList.sort()
        AlbumIDList = []
        for album in newAlbumList:
            albumID = str(self.calc_md5(album))
            barfoo = {
                "Album": album,
                "AlbumID": albumID,
            }
            AlbumIDList.append(barfoo)
        newAlbumDict = {
            "AlbumIDList": AlbumIDList,
        }
        with open(new_album_file, "w") as nf:
            data = json.dumps(newAlbumDict, indent = 4)
            nf.write(data)

        new_song_file = "/media/charliepi/FOO/music/metadata/Song_List.json"
        newSongList = list(set(self.Song_List))
        newSongList.sort()
        with open(new_song_file, "w") as nf:
            data = json.dumps(newSongList, indent = 4)
            nf.write(data)

        print(len(newArtistList))
        print(len(newAlbumList))
        print(len(newSongList))

    def main(self):
        # MF = MusicFiles()
        self.clean_metaDir()
        
        musicfiles = self.find()

        ART = Artist()
        print("\n\n STARTING ARTIST \n")
        ART.artist_check(musicfiles)
        
        print("\n STARTING ALBUM \n")
        ART.album_check(musicfiles)
        
        print("\n STARTING SONG \n")
        ART.song_check(musicfiles)

        self.get_lists()

        
 

if __name__ == "__main__":
    MF = MusicFiles()
    MF.main()
    
    
    
