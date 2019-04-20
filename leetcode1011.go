func shipWithinDays(weights []int, D int) int {

    var sum int

    var maxWeights int

    for _,v := range weights {

        sum += v

        if v > maxWeights {

            maxWeights = v

        }

    }

    if D == 1 {

        return sum

    }

    if D >= len(weights) {

        return maxWeights

    }

    var minRe int

    var maxRe int = sum

    if sum/D > maxWeights {

        minRe = sum/D

    } else {

        minRe = maxWeights

    }

    for minRe <= maxRe {

        mid := minRe + (maxRe-minRe)/2

        s, d := 0,1

        for i:=0;i<len(weights);i++ {

            s += weights[i]

            if s > mid {

                s = weights[i]

                d++

            }

        }

        if d<=D {

            maxRe = mid - 1

        } else {

            minRe = mid + 1

        }

    }

    return minRe

}