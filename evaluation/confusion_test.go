package evaluation

import (
	"math"
	"testing"
)

func TestMetrics(testEnv *testing.T) {
	confusionMat := make(ConfusionMatrix)
	confusionMat["a"] = make(map[string]int)
	confusionMat["b"] = make(map[string]int)
	confusionMat["a"]["a"] = 75
	confusionMat["a"]["b"] = 5
	confusionMat["b"]["a"] = 10
	confusionMat["b"]["b"] = 10

	tp := GetTruePositives("a", confusionMat)
	if math.Abs(tp-75) >= 1 {
		testEnv.Error(tp)
	}
	tp = GetTruePositives("b", confusionMat)
	if math.Abs(tp-10) >= 1 {
		testEnv.Error(tp)
	}

	fn := GetFalseNegatives("a", confusionMat)
	if math.Abs(fn-5) >= 1 {
		testEnv.Error(fn)
	}
	fn = GetFalseNegatives("b", confusionMat)
	if math.Abs(fn-10) >= 1 {
		testEnv.Error(fn)
	}

	tn := GetTrueNegatives("a", confusionMat)
	if math.Abs(tn-10) >= 1 {
		testEnv.Error(tn)
	}
	tn = GetTrueNegatives("b", confusionMat)
	if math.Abs(tn-75) >= 1 {
		testEnv.Error(tn)
	}

	fp := GetFalsePositives("a", confusionMat)
	if math.Abs(fp-10) >= 1 {
		testEnv.Error(fp)
	}

	fp = GetFalsePositives("b", confusionMat)
	if math.Abs(fp-5) >= 1 {
		testEnv.Error(fp)
	}

	precision := GetPrecision("a", confusionMat)
	recall := GetRecall("a", confusionMat)

	if math.Abs(precision-0.88) >= 0.01 {
		testEnv.Error(precision)
	}

	if math.Abs(recall-0.94) >= 0.01 {
		testEnv.Error(recall)
	}

	precision = GetPrecision("b", confusionMat)
	recall = GetRecall("b", confusionMat)
	if math.Abs(precision-0.666) >= 0.01 {
		testEnv.Error(precision)
	}

	if math.Abs(recall-0.50) >= 0.01 {
		testEnv.Error(recall)
	}

	precision = GetMicroPrecision(confusionMat)
	if math.Abs(precision-0.85) >= 0.01 {
		testEnv.Error(precision)
	}

	recall = GetMicroRecall(confusionMat)
	if math.Abs(recall-0.85) >= 0.01 {
		testEnv.Error(recall)
	}

	precision = GetMacroPrecision(confusionMat)
	if math.Abs(precision-0.775) >= 0.01 {
		testEnv.Error(precision)
	}

	recall = GetMacroRecall(confusionMat)
	if math.Abs(recall-0.719) > 0.01 {
		testEnv.Error(recall)
	}

	fmeasure := GetF1Score("a", confusionMat)
	if math.Abs(fmeasure-0.91) >= 0.1 {
		testEnv.Error(fmeasure)
	}

	accuracy := GetAccuracy(confusionMat)
	if math.Abs(accuracy-0.85) >= 0.1 {
		testEnv.Error(accuracy)
	}
}
