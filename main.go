package main

import (
    "fmt"
    "math/rand"
    "time"
	"os"

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

func init() {
    rootCmd.PersistentFlags().String("lang", "es", "idioma del juego (es/en)")
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
