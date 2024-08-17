package config

import "golang.org/x/crypto/bcrypt"

const AppName = "PlanetBI"
const ServerPort = "8080"
const ClientPort = "5173"
const Domain = "planet.bi"
const Sender = "PlanetBI <marawan.nwh@gmail.com>"

const HashCost = bcrypt.DefaultCost
