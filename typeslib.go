package main

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

type ArtVieW2 struct {
	Artist   string              `bson:"artist"`
	ArtistID string              `bson:"artistID"`
	Albums   []map[string]string `bson:"albums"`
	AlbCount string              `bson:"albcount"`
	Page     string              `bson:"page"`
	Index    string              `bson:"idx"`
}
type AlbVieW2 struct {
	Artist         string              `bson:"artist"`
	ArtistID       string              `bson:"artistID"`
	Album          string              `bson:"album"`
	AlbumID        string              `bson:"albumID"`
	Songs          []map[string]string `bson:"songs"`
	AlbumPage      string              `bson:"albumpage"`
	NumSongs       string              `bson:"numsongs"`
	Img_base64_str string              `bson:"img_base64_str"`
	Idx            string              `bson:"idx"`
	// PicHttpAddr string              `bson:"picHttpAddr"`
}

type AmpgoRandomPlaylistData struct {
	PlayListName  string              `bson:"playlistname"`
	PlayListID    string              `bson:"playlistID"`
	PlayListCount string              `bson:"playlistcount"`
	PlayList      []map[string]string `bson:"playlist"`
}

type JsonJPG struct {
	BaseDir        string
	Full_Filename  string
	File_Size      string
	Ext            string
	Filename       string
	Dir_catagory   string
	Dir_artist     string
	Dir_album      string
	Index          string
	Dir_delem      string
	File_id        string
	Jpg_width      string
	Jpg_height     string
	File_delem     string
	Img_base64_str string
}

type JsonMP3 struct {
	BaseDir        string
	Full_Filename  string
	File_Size      string
	Ext            string
	Dir            string
	Filename       string
	Dir_catagory   string
	Dir_artist     string
	Dir_album      string
	Dir_delem      string
	File_delem     string
	Track          string
	File_artist    string
	File_album     string
	File_song      string
	Index          string
	File_id        string
	Tags_artist    string
	Tags_album     string
	Tags_song      string
	Artist_first   string
	Album_first    string
	Song_first     string
	Img_base64_str string
	Play_length    string
}

type JsonPage struct {
	Page     string
	PageList []JsonMP3
}

type Imageinfomap struct {
	Dirpath       string `bson:"dirpath"`
	Filename      string `bson:"filename"`
	Imagesize     string `bson:"imagesize"`
	ImageHttpAddr string `bson:"imageHttpAddr"`
	Index         string `bson:"index"`
	IType         string `bson:"itype"`
	Page          string `bson:"page"`
}

type Fjpg struct {
	exists bool
	path   string
}

type randDb struct {
	PlayListName  string              `bson:"playlistname"`
	PlayListID    string              `bson:"playlistID"`
	PlayListCount string              `bson:"playlistcount"`
	Playlist      []map[string]string `bson:"playlist"`
}
