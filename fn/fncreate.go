package fn

import "fmt"

func FnCreate(flag string) string {
	return fmt.Sprintf("Le groupe %s a été crée avec succès", flag)
}