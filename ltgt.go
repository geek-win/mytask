package tpl // 指定package是tpl，我们写自定义的模板的时候，也需要指定为tpl

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }} // 使用模板的时候，会传入相应的RulesContext，Filed和Rules是其的字段，定义变量f，r，f是域，r是规则
	{{ if $r.Lt }} // 看r的Lt是否有，如果有看是否有Gt
		{{ if $r.Gt }} // 判断r是否设置了Gt这个规则
			{{  if gt $r.GetLt $r.GetGt }} // 如果小于的数字大于大于的数字，那ok
				// 通过accessor . 获取传入的值，是val，如果小于等于小的数字，大于等于小于的数字，那么有错
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					return {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLt ")" }}
					// 返回err(., "value must be inside range (" $r.GetGt ", " $r.GetLt ")")，err是在register.go中定义的函数，并且已经添加到了模板的函数映射关系中，模板可以直接调用
				}
			{{ else }} //假设此时是<a，>b，并且a<=b，那么此时val的范围应该是[a, b]之外，通过accessor .可以获取要进行验证的值
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					return {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGt "]" }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					return {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLt ")" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					return {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGte ")" }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				return {{ err . "value must be less than " $r.GetLt }}
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					return {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLte "]" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					return {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGt "]" }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					return {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLte "]" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					return {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGte ")" }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				return {{ err . "value must be less than or equal to " $r.GetLte }}
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			return {{ err . "value must be greater than " $r.GetGt }}
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			return {{ err . "value must be greater than or equal to " $r.GetGte }}
		}
	{{ end }}
`
