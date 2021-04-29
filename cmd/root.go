package cmd

import (
	"fmt"

	"github.com/odmishien/gomah/parser"
	"github.com/odmishien/gomah/rules"
	"github.com/spf13/cobra"
)

var cfgFile string

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gomah",
		Short: "gomah converts string of mahjong cards to emoji of it.",
		RunE: func(cmd *cobra.Command, args []string) error {
			var result string
			m, _ := cmd.Flags().GetString("manzu")
			rm, err := parser.GetUnicodes(m, rules.Man())
			if err != nil {
				return err
			}
			for _, um := range rm {
				result += fmt.Sprintf("%c ", um)
			}

			s, _ := cmd.Flags().GetString("souzu")
			rs, err := parser.GetUnicodes(s, rules.Sou())
			if err != nil {
				return err
			}
			for _, us := range rs {
				result += fmt.Sprintf("%c ", us)
			}

			p, _ := cmd.Flags().GetString("pinzu")
			rp, err := parser.GetUnicodes(p, rules.Pin())
			if err != nil {
				return err
			}
			for _, up := range rp {
				result += fmt.Sprintf("%c ", up)
			}

			w, _ := cmd.Flags().GetString("wind")
			rw, err := parser.GetUnicodes(w, rules.Wind())
			if err != nil {
				return err
			}
			for _, uw := range rw {
				result += fmt.Sprintf("%c ", uw)
			}

			d, _ := cmd.Flags().GetString("dragon")
			rd, err := parser.GetUnicodes(d, rules.Dragon())
			if err != nil {
				return err
			}
			for _, ud := range rd {
				result += fmt.Sprintf("%c ", ud)
			}

			cmd.Println(result)
			return nil
		},
	}

	rootCmd.Flags().StringP("manzu", "m", "", "Kind of manzu")
	rootCmd.Flags().StringP("souzu", "s", "", "Kind of souzu")
	rootCmd.Flags().StringP("pinzu", "p", "", "Kind of pinzu")
	rootCmd.Flags().StringP("wind", "w", "", "Kind of wind")
	rootCmd.Flags().StringP("dragon", "d", "", "Kind of dragon")

	return rootCmd
}
func Execute() {
	rootCmd := NewRootCmd()
	cobra.CheckErr(rootCmd.Execute())
}
