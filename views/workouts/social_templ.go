// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package workouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/jovandeginste/workout-tracker/views/helpers"
	"slices"
	"sort"
	"strings"
)

func workoutDataTitle(ctx context.Context, w *database.Workout) string {
	preferredUnits := helpers.CurrentUser(ctx).PreferredUnits()

	return helpers.I18n(
		ctx,
		`I completed a workout: %s.`,
		helpers.I18n(ctx, w.Type.String()),
	) + " " + helpers.I18n(
		ctx,
		`It took me %s to go %s. I averaged %s.`,
		helpers.HumanDuration(w.TotalDuration()),
		helpers.HumanDistance(ctx, w.TotalDistance())+" "+preferredUnits.Distance(),
		helpers.HumanSpeed(ctx, w.AverageSpeed())+" "+preferredUnits.Speed(),
	)
}
func workoutDataTags(ctx context.Context, w *database.Workout) string {
	tags := []string{
		"workout",
		helpers.I18n(ctx, w.Type.String()),
	}

	if c := w.City(); c != "" {
		tags = append(tags, c)
	}

	sort.Strings(tags)
	slices.Compact(tags)

	return strings.Join(tags, ",")
}

func Social(w *database.Workout) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"shareon print:hidden\" data-url=\" \" data-title=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(workoutDataTitle(ctx, w))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/workouts/social.templ`, Line: 47, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-hashtags=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(workoutDataTags(ctx, w))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/workouts/social.templ`, Line: 48, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><a class=\"twitter\"></a> <a class=\"mastodon\"></a> <a class=\"pinterest\"></a> <a class=\"reddit\"></a> <a class=\"teams\"></a> <a class=\"whatsapp\"></a> <a class=\"print\"></a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate