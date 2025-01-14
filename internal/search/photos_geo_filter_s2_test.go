package search

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/internal/form"
)

func TestPhotosGeoFilterS2(t *testing.T) {
	t.Run("1ef744d1e283", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "1ef744d1e283"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 3)
	})
	t.Run("85d1ea7d382c", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "85d1ea7d382c"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 8)
	})
	t.Run("StartsWithPercent", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "%gold"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterPercent", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "I love % dog"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithPercent", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "sale%"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "&IlikeFood"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterAmpersand", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Pets & Dogs"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Light&"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "'Family"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterSingleQuote", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Father's type"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Ice Cream'"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "*Forrest"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterAsterisk", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "My*Kids"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Yoga***"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithPipe", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "|Banana"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterPipe", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Red|Green"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithPipe", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Blue|"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithNumber", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "345 Shirt"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterNumber", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "type555 Blue"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithNumber", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.S2 = "Route 66"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
}

func TestPhotosGeoQueryS2(t *testing.T) {
	t.Run("s2:1ef744d1e283", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:1ef744d1e283"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 3)
	})
	t.Run("s2:85d1ea7d382c", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:85d1ea7d382c"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 8)
	})
	t.Run("85d1ea7d382c pipe 1ef744d1e283", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:85d1ea7d382c|1ef744d1e283"

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithPercent", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"%gold\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterPercent", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"I love % dog\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithPercent", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"sale%\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"&IlikeFood\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterAmpersand", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Pets & Dogs\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithAmpersand", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Light&\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"'Family\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterSingleQuote", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Father's type\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithSingleQuote", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Ice Cream'\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"*Forrest\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterAsterisk", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"My*Kids\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithAsterisk", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Yoga***\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithPipe", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"|Banana\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterPipe", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Red|Green\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithPipe", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Blue|\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("StartsWithNumber", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"345 Shirt\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("CenterNumber", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"type555 Blue\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
	t.Run("EndsWithNumber", func(t *testing.T) {
		var f form.SearchPhotosGeo

		f.Query = "s2:\"Route 66\""

		// Parse query string and filter.
		if err := f.ParseQueryString(); err != nil {
			t.Fatal(err)
		}

		photos, err := PhotosGeo(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(photos), 0)
	})
}
