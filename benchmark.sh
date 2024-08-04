copy_files() {
    local src_dir="$1"
    local dest_dir="$2"

    for item in "$src_dir"/*; do
        base_name=$(basename "$item")
        new_name="${base_name/_test/}"
        if [ -d "$item" ]; then
            mkdir -p "$dest_dir/$new_name"
            copy_files "$item" "$dest_dir/$new_name"
        else
            cp "$item" "$dest_dir/$new_name"
        fi
    done
}

destination="./.elemental"

rm -rf $destination && mkdir -p "$destination"

copy_files "./tests" "$destination"

go test -v -p 1 ./...