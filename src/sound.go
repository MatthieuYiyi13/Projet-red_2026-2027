package projet

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var ctrl *beep.Ctrl

func playMusic(path string) {
	fmt.Println("Ouverture de :", path)

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur ouverture fichier :", err)
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Erreur décodage MP3 :", err)
		f.Close()
		return
	}

	fmt.Println("MP3 décodé")

	err = speaker.Init(
		format.SampleRate,
		format.SampleRate.N(100*time.Millisecond),
	)
	if err != nil {
		fmt.Println("Erreur initialisation speaker :", err)
		streamer.Close()
		f.Close()
		return
	}

	ctrl = &beep.Ctrl{
		Streamer: beep.Loop(-1, streamer),
		Paused:   false,
	}

	speaker.Play(ctrl)

	fmt.Println("Lecture lancée")

	// Laisse le moteur audio démarrer avant le lancement du menu.
	time.Sleep(500 * time.Millisecond)
}

func PlaySoundAsyncDebut() {
	playMusic("./docs/game_of_thrones.mp3")
}