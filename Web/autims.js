const Season = 1
const Episode = 8
const Time = 40
const Pub = 0

function Calcul() {
    x = Season
    y = Episode
    z = Time
    Calc = x * y * z + Pub
    console.log(Calc + " min")
    Calc1 = Calc * 3600
    console.log(Calc1 + " secondes")
    Calc2 = Calc / 60
    console.log(Calc2 + " heures")
}
Calcul()