package structure

import (
	"fmt"
	"testing"
)

// 适配器模式
// 播放接口
type Player interface {
	PlayMusic()
}

//音乐播放器
type MusicPlayer struct {
	src string
}

func (p MusicPlayer) Playmusic() {
	fmt.Println("play music", p.src)
}

//游戏声音播放器
type GameSoundPlayer struct {
	src string
}

func (p GameSoundPlayer) PlayMusic() {
	fmt.Println("play game sound: ", p.src)
}

//游戏音乐播放器适配音乐播放器
type GameSoundPlayerAdapter struct {
	SoundPlayer GameSoundPlayer
}

func (g GameSoundPlayerAdapter) PlayMusic() {
	g.SoundPlayer.PlayMusic()
}

func TestAdapterModel(t *testing.T) {
	player := GameSoundPlayer{src: "game"}
	playerAdapter := GameSoundPlayerAdapter{SoundPlayer: player}
	play(playerAdapter)
}

func play(player Player) {
	player.PlayMusic()
}
