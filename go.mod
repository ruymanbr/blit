module github.com/ruymanbr/blit.git

go 1.16

require (
	github.com/ruymanbr/blit/pkg/blit_api v0.0.0-00010101000000-000000000000
	github.com/ruymanbr/blit/pkg/blit_cli v0.0.0-00010101000000-000000000000
)

replace github.com/ruymanbr/blit/pkg/blit_apils => ./pkg/blit_api

replace github.com/ruymanbr/blit/pkg/blit_cli => ./pkg/blit_cli
