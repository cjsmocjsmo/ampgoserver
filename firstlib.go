package main

import (
	"strings"
	"strconv"
)

func ArtistFirst(astring string) string {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
	CheckError(err, "ArtistFirst: Connections has failed")
	defer Close(client, ctx, cancel)

	char := strings.ToUpper(astring[:1])

	switch {
	case char == "A":
		var item map[string]string = map[string]string{"artist": astring}
		_, erra := InsertOne(client, ctx, "artistalpha", "A", item)
		CheckError(erra, "ArtistFirst: A insertion has failed")
		return "A Created"

	case char == "B":
		var item map[string]string = map[string]string{"artist": astring}
		_, errb := InsertOne(client, ctx, "artistalpha", "B", item)
		CheckError(errb, "ArtistFirst: B insertion has failed")
		return "B Created"

	case char == "C":
		var item map[string]string = map[string]string{"artist": astring}
		_, errc := InsertOne(client, ctx, "artistalpha", "C", item)
		CheckError(errc, "ArtistFirst: C insertion has failed")
		return "C Created"

	case char == "D":
		var item map[string]string = map[string]string{"artist": astring}
		_, errd := InsertOne(client, ctx, "artistalpha", "D", item)
		CheckError(errd, "ArtistFirst: D insertion has failed")
		return "D Created"

	case char == "E":
		var item map[string]string = map[string]string{"artist": astring}
		_, erre := InsertOne(client, ctx, "artistalpha", "E", item)
		CheckError(erre, "ArtistFirst: E insertion has failed")
		return "E Created"

	case char == "F":
		var item map[string]string = map[string]string{"artist": astring}
		_, errf := InsertOne(client, ctx, "artistalpha", "F", item)
		CheckError(errf, "ArtistFirst: F insertion has failed")
		return "F Created"

	case char == "G":
		var item map[string]string = map[string]string{"artist": astring}
		_, errg := InsertOne(client, ctx, "artistalpha", "G", item)
		CheckError(errg, "ArtistFirst: G insertion has failed")
		return "G Created"

	case char == "H":
		var item map[string]string = map[string]string{"artist": astring}
		_, errh := InsertOne(client, ctx, "artistalpha", "H", item)
		CheckError(errh, "ArtistFirst: H insertion has failed")
		return "H Created"

	case char == "I":
		var item map[string]string = map[string]string{"artist": astring}
		_, erri := InsertOne(client, ctx, "artistalpha", "I", item)
		CheckError(erri, "ArtistFirst: I insertion has failed")
		return "I Created"

	case char == "J":
		var item map[string]string = map[string]string{"artist": astring}
		_, errj := InsertOne(client, ctx, "artistalpha", "J", item)
		CheckError(errj, "ArtistFirst: J insertion has failed")
		return "J Created"

	case char == "K":
		var item map[string]string = map[string]string{"artist": astring}
		_, errk := InsertOne(client, ctx, "artistalpha", "K", item)
		CheckError(errk, "ArtistFirst: K insertion has failed")
		return "K Created"

	case char == "L":
		var item map[string]string = map[string]string{"artist": astring}
		_, errl := InsertOne(client, ctx, "artistalpha", "L", item)
		CheckError(errl, "ArtistFirst: L insertion has failed")
		return "L Created"

	case char == "M":
		var item map[string]string = map[string]string{"artist": astring}
		_, errm := InsertOne(client, ctx, "artistalpha", "M", item)
		CheckError(errm, "ArtistFirst: M insertion has failed")
		return "M Created"

	case char == "N":
		var item map[string]string = map[string]string{"artist": astring}
		_, errn := InsertOne(client, ctx, "artistalpha", "N", item)
		CheckError(errn, "ArtistFirst: N insertion has failed")
		return "N Created"

	case char == "O":
		var item map[string]string = map[string]string{"artist": astring}
		_, erro := InsertOne(client, ctx, "artistalpha", "O", item)
		CheckError(erro, "ArtistFirst: O insertion has failed")
		return "O Created"

	case char == "P":
		var item map[string]string = map[string]string{"artist": astring}
		_, errp := InsertOne(client, ctx, "artistalpha", "P", item)
		CheckError(errp, "ArtistFirst: P insertion has failed")
		return "P Created"

	case char == "Q":
		var item map[string]string = map[string]string{"artist": astring}
		_, errq := InsertOne(client, ctx, "artistalpha", "Q", item)
		CheckError(errq, "ArtistFirst: Q insertion has failed")
		return "Q Created"

	case char == "R":
		var item map[string]string = map[string]string{"artist": astring}
		_, errr := InsertOne(client, ctx, "artistalpha", "R", item)
		CheckError(errr, "ArtistFirst: R insertion has failed")
		return "R Created"

	case char == "S":
		var item map[string]string = map[string]string{"artist": astring}
		_, errs := InsertOne(client, ctx, "artistalpha", "S", item)
		CheckError(errs, "ArtistFirst: S insertion has failed")
		return "S Created"

	case char == "T":
		var item map[string]string = map[string]string{"artist": astring}
		_, errt := InsertOne(client, ctx, "artistalpha", "T", item)
		CheckError(errt, "ArtistFirst: T insertion has failed")
		return "T Created"

	case char == "U":
		var item map[string]string = map[string]string{"artist": astring}
		_, erru := InsertOne(client, ctx, "artistalpha", "U", item)
		CheckError(erru, "ArtistFirst: U insertion has failed")
		return "U Created"

	case char == "V":
		var item map[string]string = map[string]string{"artist": astring}
		_, errv := InsertOne(client, ctx, "artistalpha", "V", item)
		CheckError(errv, "ArtistFirst: V insertion has failed")
		return "V Created"

	case char == "W":
		var item map[string]string = map[string]string{"artist": astring}
		_, errw := InsertOne(client, ctx, "artistalpha", "W", item)
		CheckError(errw, "ArtistFirst: W insertion has failed")
		return "W Created"

	case char == "X":
		var item map[string]string = map[string]string{"artist": astring}
		_, errx := InsertOne(client, ctx, "artistalpha", "X", item)
		CheckError(errx, "ArtistFirst: X insertion has failed")
		return "X Created"

	case char == "Z":
		var item map[string]string = map[string]string{"artist": astring}
		_, errz := InsertOne(client, ctx, "artistalpha", "Z", item)
		CheckError(errz, "ArtistFirst: Z insertion has failed")
		return "Z Created"
	}
	return "None"
}

func AlbumFirst(astring string) string {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
	CheckError(err, "AlbumFirst:  Connections has failed")
	defer Close(client, ctx, cancel)

	// char := StartsWith(astring)

	char := strings.ToUpper(astring[:1])

	switch {
	case char == "A":
		var item map[string]string = map[string]string{"album": astring}
		_, erra := InsertOne(client, ctx, "albumalpha", "A", item)
		CheckError(erra, "AlbumFirst: A insertion has failed")
		return "A Created"

	case char == "B":
		var item map[string]string = map[string]string{"album": astring}
		_, errb := InsertOne(client, ctx, "albumalpha", "B", item)
		CheckError(errb, "AlbumFirst: B insertion has failed")
		return "B Created"

	case char == "C":
		var item map[string]string = map[string]string{"album": astring}
		_, errc := InsertOne(client, ctx, "albumalpha", "C", item)
		CheckError(errc, "AlbumFirst: C insertion has failed")
		return "C Created"

	case char == "D":
		var item map[string]string = map[string]string{"album": astring}
		_, errd := InsertOne(client, ctx, "albumalpha", "D", item)
		CheckError(errd, "AlbumFirst: D insertion has failed")
		return "D Created"

	case char == "E":
		var item map[string]string = map[string]string{"album": astring}
		_, erre := InsertOne(client, ctx, "albumalpha", "E", item)
		CheckError(erre, "AlbumFirst: E insertion has failed")
		return "E Created"

	case char == "F":
		var item map[string]string = map[string]string{"album": astring}
		_, errf := InsertOne(client, ctx, "albumalpha", "F", item)
		CheckError(errf, "AlbumFirst: F insertion has failed")
		return "F Created"

	case char == "G":
		var item map[string]string = map[string]string{"album": astring}
		_, errg := InsertOne(client, ctx, "albumalpha", "G", item)
		CheckError(errg, "AlbumFirst: G insertion has failed")
		return "G Created"

	case char == "H":
		var item map[string]string = map[string]string{"album": astring}
		_, errh := InsertOne(client, ctx, "albumalpha", "H", item)
		CheckError(errh, "AlbumFirst: H insertion has failed")
		return "H Created"

	case char == "I":
		var item map[string]string = map[string]string{"album": astring}
		_, erri := InsertOne(client, ctx, "albumalpha", "I", item)
		CheckError(erri, "AlbumFirst: I insertion has failed")
		return "I Created"

	case char == "J":
		var item map[string]string = map[string]string{"album": astring}
		_, errj := InsertOne(client, ctx, "albumalpha", "J", item)
		CheckError(errj, "AlbumFirst: J insertion has failed")
		return "J Created"

	case char == "K":
		var item map[string]string = map[string]string{"album": astring}
		_, errk := InsertOne(client, ctx, "albumalpha", "K", item)
		CheckError(errk, "AlbumFirst: K insertion has failed")
		return "K Created"

	case char == "L":
		var item map[string]string = map[string]string{"album": astring}
		_, errl := InsertOne(client, ctx, "albumalpha", "L", item)
		CheckError(errl, "AlbumFirst: L insertion has failed")
		return "L Created"

	case char == "M":
		var item map[string]string = map[string]string{"album": astring}
		_, errm := InsertOne(client, ctx, "albumalpha", "M", item)
		CheckError(errm, "AlbumFirst: M insertion has failed")
		return "M Created"

	case char == "N":
		var item map[string]string = map[string]string{"album": astring}
		_, errn := InsertOne(client, ctx, "albumalpha", "N", item)
		CheckError(errn, "AlbumFirst: N insertion has failed")
		return "N Created"

	case char == "O":
		var item map[string]string = map[string]string{"album": astring}
		_, erro := InsertOne(client, ctx, "albumalpha", "O", item)
		CheckError(erro, "AlbumFirst: O insertion has failed")
		return "O Created"

	case char == "P":
		var item map[string]string = map[string]string{"album": astring}
		_, errp := InsertOne(client, ctx, "albumalpha", "P", item)
		CheckError(errp, "AlbumFirst: P insertion has failed")
		return "P Created"

	case char == "Q":
		var item map[string]string = map[string]string{"album": astring}
		_, errq := InsertOne(client, ctx, "albumalpha", "Q", item)
		CheckError(errq, "AlbumFirst: Q insertion has failed")
		return "Q Created"

	case char == "R":
		var item map[string]string = map[string]string{"album": astring}
		_, errr := InsertOne(client, ctx, "albumalpha", "R", item)
		CheckError(errr, "AlbumFirst: R insertion has failed")
		return "R Created"

	case char == "S":
		var item map[string]string = map[string]string{"album": astring}
		_, errs := InsertOne(client, ctx, "albumalpha", "S", item)
		CheckError(errs, "AlbumFirst: S insertion has failed")
		return "S Created"

	case char == "T":
		var item map[string]string = map[string]string{"album": astring}
		_, errt := InsertOne(client, ctx, "albumalpha", "T", item)
		CheckError(errt, "AlbumFirst: T insertion has failed")
		return "T Created"

	case char == "U":
		var item map[string]string = map[string]string{"album": astring}
		_, erru := InsertOne(client, ctx, "albumalpha", "U", item)
		CheckError(erru, "AlbumFirst: U insertion has failed")
		return "U Created"

	case char == "V":
		var item map[string]string = map[string]string{"album": astring}
		_, errv := InsertOne(client, ctx, "albumalpha", "V", item)
		CheckError(errv, "AlbumFirst: V insertion has failed")
		return "V Created"

	case char == "W":
		var item map[string]string = map[string]string{"album": astring}
		_, errw := InsertOne(client, ctx, "albumalpha", "W", item)
		CheckError(errw, "AlbumFirst: W insertion has failed")
		return "W Created"

	case char == "X":
		var item map[string]string = map[string]string{"album": astring}
		_, errx := InsertOne(client, ctx, "albumalpha", "X", item)
		CheckError(errx, "AlbumFirst: X insertion has failed")
		return "X Created"

	case char == "Z":
		var item map[string]string = map[string]string{"album": astring}
		_, errz := InsertOne(client, ctx, "albumalpha", "Z", item)
		CheckError(errz, "AlbumFirst: Z insertion has failed")
		return "Z Created"
	}
	return "None"
}

func SongFirst() string {

	aAll := AmpgoFind("maindb", "maindb", "Song_first", "A")
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
	CheckError(err, "SongFirst: Connections has failed")
	defer Close(client, ctx, cancel)
	for _, a := range aAll {
		_, err = InsertOne(client, ctx, "songalpha", "A", a)
		CheckError(err, "SongFirst: a insertion has failed")
	}
	aa := len(aAll)

	bAll := AmpgoFind("maindb", "maindb", "Song_first", "B")
	for _, b := range bAll {
		_, err = InsertOne(client, ctx, "songalpha", "B", b)
		CheckError(err, "SongFirst: b insertion has failed")
	}
	bb := len(bAll)

	cAll := AmpgoFind("maindb", "maindb", "Song_first", "C")
	for _, c := range cAll {
		_, err = InsertOne(client, ctx, "songalpha", "C", c)
		CheckError(err, "SongFirst: c insertion has failed")
	}
	cc := len(cAll)

	dAll := AmpgoFind("maindb", "maindb", "Song_first", "D")
	for _, d := range dAll {
		_, err = InsertOne(client, ctx, "songalpha", "D", d)
		CheckError(err, "SongFirst: d insertion has failed")
	}
	dd := len(dAll)

	eAll := AmpgoFind("maindb", "maindb", "Song_first", "E")
	for _, e := range eAll {
		_, err = InsertOne(client, ctx, "songalpha", "E", e)
		CheckError(err, "SongFirst: e insertion has failed")
	}
	ee := len(eAll)

	fAll := AmpgoFind("maindb", "maindb", "Song_first", "F")
	for _, f := range fAll {
		_, err = InsertOne(client, ctx, "songalpha", "F", f)
		CheckError(err, "SongFirst: f insertion has failed")
	}
	ff := len(fAll)

	gAll := AmpgoFind("maindb", "maindb", "Song_first", "G")
	gg := len(gAll)
	for _, g := range gAll {
		_, err = InsertOne(client, ctx, "songalpha", "G", g)
		CheckError(err, "SongFirst: g insertion has failed")
	}

	hAll := AmpgoFind("maindb", "maindb", "Song_first", "H")
	hh := len(hAll)
	for _, h := range hAll {
		_, err = InsertOne(client, ctx, "songalpha", "H", h)
		CheckError(err, "SongFirst: h insertion has failed")
	}

	iAll := AmpgoFind("maindb", "maindb", "Song_first", "I")
	ii := len(iAll)
	for _, i := range iAll {
		_, err = InsertOne(client, ctx, "songalpha", "I", i)
		CheckError(err, "SongFirst: i insertion has failed")
	}

	jAll := AmpgoFind("maindb", "maindb", "Song_first", "J")
	jj := len(jAll)
	for _, j := range jAll {
		_, err = InsertOne(client, ctx, "songalpha", "J", j)
		CheckError(err, "SongFirst: j insertion has failed")
	}

	kAll := AmpgoFind("maindb", "maindb", "Song_first", "K")
	kk := len(kAll)
	for _, k := range kAll {
		_, err = InsertOne(client, ctx, "songalpha", "K", k)
		CheckError(err, "SongFirst: k insertion has failed")
	}

	lAll := AmpgoFind("maindb", "maindb", "Song_first", "L")
	ll := len(lAll)
	for _, l := range lAll {
		_, err = InsertOne(client, ctx, "songalpha", "L", l)
		CheckError(err, "SongFirst: l insertion has failed")
	}

	mAll := AmpgoFind("maindb", "maindb", "Song_first", "M")
	mm := len(mAll)
	for _, m := range mAll {
		_, err = InsertOne(client, ctx, "songalpha", "M", m)
		CheckError(err, "SongFirst: m insertion has failed")
	}

	nAll := AmpgoFind("maindb", "maindb", "Song_first", "N")
	nn := len(nAll)
	for _, n := range nAll {
		_, err = InsertOne(client, ctx, "songalpha", "N", n)
		CheckError(err, "SongFirst: n insertion has failed")
	}

	oAll := AmpgoFind("maindb", "maindb", "Song_first", "O")
	oo := len(oAll)
	for _, o := range oAll {
		_, err = InsertOne(client, ctx, "songalpha", "O", o)
		CheckError(err, "SongFirst: o insertion has failed")
	}

	pAll := AmpgoFind("maindb", "maindb", "Song_first", "P")
	pp := len(pAll)
	for _, p := range pAll {
		_, err = InsertOne(client, ctx, "songalpha", "P", p)
		CheckError(err, "SongFirst: p insertion has failed")
	}

	qAll := AmpgoFind("maindb", "maindb", "Song_first", "Q")
	qq := len(qAll)
	for _, q := range qAll {
		_, err = InsertOne(client, ctx, "songalpha", "Q", q)
		CheckError(err, "SongFirst: q insertion has failed")
	}

	rAll := AmpgoFind("maindb", "maindb", "Song_first", "R")
	rr := len(rAll)
	for _, r := range rAll {
		_, err = InsertOne(client, ctx, "songalpha", "R", r)
		CheckError(err, "SongFirst: r insertion has failed")
	}

	sAll := AmpgoFind("maindb", "maindb", "Song_first", "S")
	ss := len(sAll)
	for _, s := range sAll {
		_, err = InsertOne(client, ctx, "songalpha", "S", s)
		CheckError(err, "SongFirst: s insertion has failed")
	}

	tAll := AmpgoFind("maindb", "maindb", "Song_first", "T")
	tt := len(tAll)
	for _, t := range tAll {
		_, err = InsertOne(client, ctx, "songalpha", "T", t)
		CheckError(err, "SongFirst: t insertion has failed")
	}

	uAll := AmpgoFind("maindb", "maindb", "Song_first", "U")
	uu := len(uAll)
	for _, u := range uAll {
		_, err = InsertOne(client, ctx, "songalpha", "U", u)
		CheckError(err, "SongFirst: u insertion has failed")
	}

	vAll := AmpgoFind("maindb", "maindb", "Song_first", "V")
	vv := len(vAll)
	for _, v := range vAll {
		_, err = InsertOne(client, ctx, "songalpha", "V", v)
		CheckError(err, "SongFirst: v insertion has failed")
	}

	wAll := AmpgoFind("maindb", "maindb", "Song_first", "W")
	ww := len(wAll)
	for _, w := range wAll {
		_, err = InsertOne(client, ctx, "songalpha", "W", w)
		CheckError(err, "SongFirst: w insertion has failed")
	}

	xAll := AmpgoFind("maindb", "maindb", "Song_first", "X")
	xx := len(xAll)
	for _, x := range xAll {
		_, err = InsertOne(client, ctx, "songalpha", "X", x)
		CheckError(err, "SongFirst: x insertion has failed")
	}

	yAll := AmpgoFind("maindb", "maindb", "Song_first", "Y")
	yy := len(yAll)
	for _, y := range yAll {
		_, err = InsertOne(client, ctx, "songalpha", "Y", y)
		CheckError(err, "SongFirst: y insertion has failed")
	}

	zAll := AmpgoFind("maindb", "maindb", "Song_first", "Z")
	zz := len(zAll)
	for _, z := range zAll {
		_, err = InsertOne(client, ctx, "songalpha", "Z", z)
		CheckError(err, "SongFirst: z insertion has failed")
	}

	t1 := aa + bb + cc + dd + ee + ff + gg + hh + ii + jj + kk + ll + mm
	t2 := nn + oo + pp + qq + rr + ss + tt + uu + vv + ww + xx + yy + zz
	tot := t1 + t2
	total := strconv.Itoa(tot)
	var total2 map[string]string = map[string]string{"total": total}

	client, ctx, cancel, err = Connect("mongodb://db:27017/ampgo")
	CheckError(err, "SongFirst: Connections has failed")
	defer Close(client, ctx, cancel)

	_, err = InsertOne(client, ctx, "songtotal", "total", &total2)
	CheckError(err, "SongFirst: z insertion has failed")

	return "Complete"
}