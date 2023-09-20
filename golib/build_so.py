import os

command = "go build -buildmode=c-shared -o golib/golib.so"

def is_go_file(file_name: str):
    if len(file_name) < 4:
        return False
    return file_name[len(file_name)-2:] == "go"


gofiles = list(filter(is_go_file, os.listdir("./golib")))

golib = "golib"

for gofile in gofiles:
    command += " " + os.path.join(golib, gofile)

os.system(command)