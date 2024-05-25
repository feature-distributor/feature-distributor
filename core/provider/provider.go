package provider

import (
	"context"
	"encoding/json"
	"feature-distributor/core/db"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"fmt"
	"github.com/open-feature/go-sdk/openfeature"
	"github.com/pkg/errors"
	"strconv"
)

type CoreProvider struct {
	project model.Project
}

func (i CoreProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "feature-distributor",
	}
}

func (i CoreProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}

// resolutionError 逻辑出错时的详情信息
func resolutionError(err error) openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		ResolutionError: openfeature.NewGeneralResolutionError(fmt.Sprintf("failed to get value, %v", err)),
		Reason:          openfeature.UnknownReason,
	}
}

// resolutionStatic 没有动态规则时的详情信息
func resolutionStatic() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason: openfeature.StaticReason,
	}
}

// resolutionDynamic 动态规则计算之后返回的详情信息
func resolutionDynamic() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason: openfeature.TargetingMatchReason,
	}
}

// resolutionDisable 禁用规则后返回的详情信息
func resolutionDisable() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason: openfeature.DisabledReason,
	}
}

func genericResolve[T comparable](ctx context.Context, projectId int64, toggleKey string, defaultValue T, m func(string) (T, error)) (T, openfeature.ProviderResolutionDetail) {
	toggle, toggleValues, err := db.SelectToggleValues(ctx, projectId, toggleKey)
	if err != nil {
		return defaultValue, resolutionError(errors.Wrap(err, "select toggle value failed"))
	}
	valueMap := make(map[int64]T)
	for _, value := range toggleValues {
		vv, err := m(value.Value)
		if err != nil {
			return defaultValue, resolutionError(errors.Wrap(err, fmt.Sprintf("convert value failed with %s", value.Value)))
		}
		valueMap[value.ID] = vv
	}
	if !toggle.Enable {
		return valueMap[toggle.ReturnValueWhenDisable], resolutionDisable()
	}
	return valueMap[toggle.ReturnValueWhenDisable], resolutionStatic()
}

func (i CoreProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	v, r := genericResolve(ctx, i.project.ID, flag, defaultValue, func(s string) (bool, error) {
		switch s {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return false, errors.New("unknown bool value")
		}
	})
	return openfeature.BoolResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	v, r := genericResolve(ctx, i.project.ID, flag, defaultValue, func(s string) (string, error) {
		return s, nil
	})
	return openfeature.StringResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	v, r := genericResolve(ctx, i.project.ID, flag, defaultValue, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
	return openfeature.FloatResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	v, r := genericResolve(ctx, i.project.ID, flag, defaultValue, func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})
	return openfeature.IntResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	v, r := genericResolve(ctx, i.project.ID, flag, defaultValue, func(s string) (interface{}, error) {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(s), &m)
		return m, err
	})
	return openfeature.InterfaceResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) Init(evaluationContext openfeature.EvaluationContext) error {
	p := query.Project
	project, err := p.Where(p.Key.Eq(evaluationContext.TargetingKey())).First()
	if err != nil {
		return err
	}
	i.project = *project
	return nil
}

func (i CoreProvider) Status() openfeature.State {
	return openfeature.ReadyState
}

func (i CoreProvider) Shutdown() {
}
