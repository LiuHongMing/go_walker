/*
  音乐播放
 */

package player

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &Mp3Player{}
	//case "WAV":
	//	p = &WAVPlayer{}
	}

	p.Play(source)
}
