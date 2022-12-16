#!/usr/bin/env python3

import os
import re
import json
from pprint import pprint

class SortMusicFiles:

    def __init__(self) -> None:
        self.media = "/media/charliepi/FOO/music"  #os.getenv("AMPGO_MEDIA_PATH") + "/"
        self.meta = "/media/charliepi/FOO/music/metadata/" #os.getenv("AMPGO_MEDIA_METADATA_PATH") + "/"

    def find(self, apath):
        mp3_files_list = []
        jpg_files_list = []
        jpg = re.compile("jpg")
        mp3 = re.compile("mp3")
        for (paths, Dirs, files) in os.walk(apath, followlinks=True):
            for Filename in files:
                fnn = os.path.join(paths, Filename)
                if re.search(jpg, fnn) != None:
                    jpg_files_list.append(fnn)
                if re.search(mp3, fnn) != None:
                    mp3_files_list.append(fnn)
        return (jpg_files_list, mp3_files_list)

    def chunck_list(self, alist, chunck_size):
        chuncked_list = []
        for i in range(0, len(alist), chunck_size):
            chuncked_list.append(alist[i:i+chunck_size])
        return chuncked_list

    def main(self):
        jpglist, mp3list = smf.find(self.meta)
        chuncks = smf.chunck_list(mp3list, 35)
        page = 1
        for chunck in chuncks:
            newlist = []
            page += 1
            newfile = self.meta + "songs_page" + str(page) + ".json"
            for addr in chunck:
                with open(addr, "rb") as stuff:
                    data = json.load(stuff)
                    newlist.append(data)
            newDict = {
                'Page': str(page),
                'PageList': newlist,
            }
            # print(newDict)
            with open(newfile, "w") as nf:
                data2 = json.dumps(newDict, indent = 4)
                nf.write(data2)
        print("COMPLETE")


        
        




if __name__ == "__main__":
    from create_json import MusicFiles
    MF = MusicFiles()
    MF.main()
    # media = "/media/charliepi/FOO/music"  #os.getenv("AMPGO_MEDIA_PATH") + "/"
    # meta = "/home/charliepi/Music/" #os.getenv("AMPGO_MEDIA_METADATA_PATH") + "/"
    smf = SortMusicFiles()
    smf.main()