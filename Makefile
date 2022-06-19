MAKEFLAGS += -rR
OUTDIR := ./out

.PHONY : build

build : list_tracks get_track

vet :
	go vet ./...

configure :
	mkdir -p $(OUTDIR)
	go mod tidy

list_tracks : vet configure
	go build -o $(OUTDIR)/list_tracks ./handler/list_tracks/main.go

get_track : configure
	go build -o $(OUTDIR)/get_track ./handler/get_track/main.go

clean :
	rm -rf $(OUTDIR)
