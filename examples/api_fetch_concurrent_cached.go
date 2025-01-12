package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// PokemonIndex representa la respuesta del índice de Pokémon.
type PokemonIndex struct {
	Results []PokemonEntry `json:"results"`
}

type PokemonEntry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonDetails representa los detalles individuales de un Pokémon.
type PokemonDetails struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}

// Cache para almacenar respuestas ya obtenidas.
var cache sync.Map

// fetchIndex obtiene el índice de Pokémon desde el endpoint.
func fetchIndex(url string) ([]PokemonEntry, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var index PokemonIndex
	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return nil, err
	}

	return index.Results, nil
}

// fetchDetails obtiene los detalles de un Pokémon desde su URL o de la caché si ya existe.
func fetchDetails(url string, wg *sync.WaitGroup, results chan<- PokemonDetails, errors chan<- error) {
	defer wg.Done()

	// Verificar si ya está en la caché
	if value, ok := cache.Load(url); ok {
		results <- value.(PokemonDetails)
		return
	}

	// Si no está en caché, hacer la petición
	resp, err := http.Get(url)
	if err != nil {
		errors <- err
		return
	}
	defer resp.Body.Close()

	var details PokemonDetails
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		errors <- err
		return
	}

	// Guardar en caché
	cache.Store(url, details)

	results <- details
}

func main() {
	baseURL := "https://pokeapi.co/api/v2/pokemon"

	// Paso 1: Obtener el índice de Pokémon
	pokemonEntries, err := fetchIndex(baseURL)
	if err != nil {
		fmt.Printf("Error al obtener el índice: %v\n", err)
		return
	}

	// Canales para manejar resultados y errores
	results := make(chan PokemonDetails, len(pokemonEntries))
	errors := make(chan error, len(pokemonEntries))

	var wg sync.WaitGroup

	// Paso 2: Obtener los detalles de cada Pokémon de forma concurrente con caché
	for _, entry := range pokemonEntries {
		wg.Add(1)
		go fetchDetails(entry.URL, &wg, results, errors)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Cerrar los canales
	close(results)
	close(errors)

	// Mostrar resultados
	fmt.Println("Detalles obtenidos:")
	for result := range results {
		fmt.Printf("- Nombre: %s, Altura: %d, Peso: %d\n", result.Name, result.Height, result.Weight)
	}

	// Mostrar errores (si los hay)
	for err := range errors {
		fmt.Printf("Error: %v\n", err)
	}
}
