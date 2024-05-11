/*
Copyright © 2024 badonix <nikdanelia@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

// shortenCmd represents the shorten command
var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Command to shorten your url",
	Long:  `Drop in your url and i will shorten it, using the ulvis.net, you can also provide custom name`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")
		name, _ := cmd.Flags().GetString("name")
		finalUrl := fmt.Sprintf("https://ulvis.net/api.php?url=%s&custom=%s&private=1", url, name)
		if url == "" {
			return errors.New("URL is required")
		}
		if !isValidUrl(url) {
			return errors.New("URL is not valid")
		}
		resp, err := http.Get(finalUrl)
		if err != nil {
			return errors.New("some error occured")
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
	shortenCmd.Flags().StringP("url", "u", "", "URL which you want to shorten, https://example.com *REQUIRED*")
	shortenCmd.Flags().StringP("name", "n", "", "custom name for your url")
}

func isValidUrl(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return false
	}
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}
