package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "verdades",
	Short: "Un juego de verdades en Go",
	Long:  `Un juego simple de verdades para jugar con amigos.`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().Unix())
		fmt.Println("Bienvenido a Verdades en Go!")
		fmt.Println("¿Listo para jugar?")
		fmt.Println("¡Comencemos!")
		fmt.Println("¿Qué quieres hacer?")

		opciones := []string{"Verdad", "Desafío"}
		opcion := opciones[rand.Intn(len(opciones))]
		fmt.Printf("%s\n", opcion)
	},
}

var langCmd = &cobra.Command{
	Use:   "lang",
	Short: "Cambiar el idioma del juego",
	Long:  `Cambiar el idioma del juego entre español (es) e inglés (en).`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		switch lang {
		case "es":
			fmt.Println("Idioma cambiado a español.")
		case "en":
			fmt.Println("Language changed to English.")
		default:
			fmt.Println("Idioma no válido. Usando idioma predeterminado (español).")
		}
	},
}

var duracionCmd = &cobra.Command{
	Use:   "duracion",
	Short: "Configurar la duración del juego",
	Long:  `Configurar la duración del juego en minutos.`,
	Run: func(cmd *cobra.Command, args []string) {
		duracion, _ := cmd.Flags().GetInt("duracion")
		fmt.Println("La duración del juego es de", duracion, "minutos.")
	},
}

func init() {
	rootCmd.PersistentFlags().String("lang", "es", "idioma del juego (es/en)")

	rootCmd.AddCommand(langCmd)
	langCmd.PersistentFlags().String("lang", "es", "Cambiar el idioma del juego (es/en)")

	rootCmd.AddCommand(duracionCmd)
	duracionCmd.PersistentFlags().Int("duracion", 10, "Configurar la duración del juego (minutos)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Resultado por consola

// next@next-System:~/Vídeos/GO/app-go-IA$ ./mi-juego lang --lang en
// Language changed to English.
// next@next-System:~/Vídeos/GO/app-go-IA$ ./mi-juego duracion --duracion 15
// La duración del juego es de 15 minutos.
// next@next-System:~/Vídeos/GO/app-go-IA$
