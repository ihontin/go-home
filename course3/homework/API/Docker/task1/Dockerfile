FROM golang:1.19

RUN mkdir -p /home/i/work/src/studentgit.kata.academy/Alkolex/go-kata/course3/homework/API/Docker/task1/app
WORKDIR /home/i/work/src/studentgit.kata.academy/Alkolex/go-kata/course3/homework/API/Docker/task1/app

COPY . /home/i/work/src/studentgit.kata.academy/Alkolex/go-kata/course3/homework/API/Docker/task1/app
# CMD - используется для указания команды, которая будет выполнена при запуске контейнера на основе образа.
# ENTRYPOINT -
CMD ["go","run","main.go"]