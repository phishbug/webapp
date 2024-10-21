package seo

import (
    "regexp"
    "strings"
)

// GenerateSlug creates a web-friendly URL slug from a given title
func GenerateSlug(title string) string {
    // Convert to lower case
    slug := strings.ToLower(title)

    // Remove any special characters (keeping only alphanumeric and spaces)
    re := regexp.MustCompile(`[^\w\s-]`)
    slug = re.ReplaceAllString(slug, "")

    // Replace spaces and multiple dashes with a single dash
    slug = regexp.MustCompile(`[\s-]+`).ReplaceAllString(slug, "-")

    // Trim any dashes from the beginning and end
    slug = strings.Trim(slug, "-")

    return slug
}