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

type ARID struct {
	Artist string `bson:"Artist"`
	ArtistID string `bson:"ArtistID"`
}

type ArtID struct {
	ArtistIDList []ARID `bson:"ArtistIDList"`
}

type ALID struct {
	Album string `bson:"Album"`
	AlbumID string `bson:"AlbumID"`
}

type AlbID struct {
	AlbumIDList []ALID `bson:"AlbumIDList"`
}

type JsonJPG struct {
	BaseDir        string `bson:"BaseDir"`
	Full_Filename  string `bson:"Full_Filename"`
	File_Size      string `bson:"File_Size"`
	Ext            string `bson:"Ext"`
	Filename       string `bson:"Filename"`
	Dir_catagory   string `bson:"Dir_catagory"`
	Dir_artist     string `bson:"Dir_artist"`
	Dir_album      string `bson:"Dir_album"`
	Index          string `bson:"Index"`
	Dir_delem      string `bson:"Dir_delem"`
	File_id        string `bson:"File_id"`
	Jpg_width      string `bson:"Jpg_width"`
	Jpg_height     string `bson:"Jpg_height"`
	File_delem     string `bson:"File_delem"`
	Img_base64_str string `bson:"Img_base64_str"`
}

type JsonMP3 struct {
	BaseDir        string `bson:"BaseDir"`
	Full_Filename  string `bson:"Full_Filename"`
	File_Size      string `bson:"File_Size"`
	Ext            string `bson:"Ext"`
	Dir            string `bson:"Dir"`
	Filename       string `bson:"Filename"`
	Dir_catagory   string `bson:"Dir_catagory"`
	Dir_artist     string `bson:"Dir_artist"`
	Dir_album      string `bson:"Dir_album"`
	Dir_delem      string `bson:"Dir_delem"`
	File_delem     string `bson:"File_delem"`
	Track          string `bson:"Track"`
	File_artist    string `bson:"File_artist"`
	File_album     string `bson:"File_album"`
	File_song      string `bson:"File_song"`
	Index          string `bson:"Index"`
	File_id        string `bson:"File_id"`
	Tags_artist    string `bson:"Tags_artist"`
	Tags_album     string `bson:"Tags_album"`
	Tags_song      string `bson:"Tags_song"`
	Artist_first   string `bson:"Artist_first"`
	Album_first    string `bson:"Album_first"`
	Song_first     string `bson:"Song_first"`
	Img_base64_str string `bson:"Img_base64_str"`
	Play_length    string `bson:"Play_length"`
}

type JsonPage struct {
	Page     string    `bson:"Page"`
	PageList []JsonMP3 `bson:"PageList"`
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
