# new-go-sro-db
This is a mysql database containing more tables than the original `go-sro-db` did. I have copied it from the [eSRO](https://github.com/ghostuser846/eSRO/) repository and fixed some spelling and encoding issues.

To setup this database, you need to extract the Media.pk2 file into this directory. I used Veykrils [pk2_mate](https://github.com/Veykril/pk2) for this. Then run the `setup-media-pk2.sh` script. It will copy and format the textdata from the now unpacked Media directory and insert their content into the tables when starting the docker container.

You can run it with `docker compose up new_db` from this project's root directory.
