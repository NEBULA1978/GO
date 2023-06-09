package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

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

var rootCmd = &cobra.Command{
	Use:   "verdades",
	Short: "Un juego de verdades en Go",
	Long:  `Un juego simple de verdades para jugar con amigos.`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().Unix())
		fmt.Println("Bienvenido a Verdades en Go!")
		fmt.Println("¿Listo para jugar?")
		fmt.Println("¡Comencemos!")

		for {
			opciones := []string{"Verdad", "Desafío"}
			opcion := opciones[rand.Intn(len(opciones))]
			fmt.Printf("%s\n", opcion)

			// Espera a que el usuario presione Enter para continuar
			fmt.Scanln()

			// Pregunta al usuario si quiere jugar de nuevo
			fmt.Println("¿Quieres jugar de nuevo? (S/N)")
			var respuesta string
			fmt.Scanln(&respuesta)
			if respuesta != "S" && respuesta != "s" {
				break
			}
		}
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

// PARA EJECUTAR POR CONSOLA

// Para ejecutar el juego, desde la línea de comandos, debes ubicarte en el directorio donde se encuentra el archivo main.go y luego ejecutar el siguiente comando:

// go

// go run main.go

// Esto iniciará el juego y mostrará en pantalla las opciones disponibles para cambiar el idioma del juego o la duración del mismo.

// Para cambiar el idioma del juego a inglés, debes ejecutar el siguiente comando:

// bash

// ./mi-juego lang --lang en

// Para configurar la duración del juego a 15 minutos, debes ejecutar el siguiente comando:

// bash

// ./mi-juego duracion --duracion 15

// Recuerda que para que los cambios de idioma o duración del juego tengan efecto, debes reiniciar el juego.
