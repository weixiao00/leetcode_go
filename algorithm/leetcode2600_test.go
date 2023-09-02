package algorithm

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	} else if k <= numOnes+numZeros {
		return numOnes
	} else {
		return numOnes - (k - numOnes - numZeros)
	}
}
