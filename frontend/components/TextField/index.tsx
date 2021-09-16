import React, { ChangeEvent, ReactElement, Ref, useState } from "react"
import { useId } from "../../lib/useId"

export default function TextField(props: React.DetailedHTMLProps<React.InputHTMLAttributes<HTMLInputElement>, HTMLInputElement> & {
    light?: boolean,
    containerProps?: React.DetailedHTMLProps<React.HTMLAttributes<HTMLDivElement>, HTMLDivElement>
}) {
    const [sFilled, setFilled] = useState(false)
    const [id, idRef] = useId()

    const filled = (props.value != null) ? props.value.toString().length > 0 : sFilled

    const onChange = (e: ChangeEvent<HTMLInputElement>) => {
        setFilled(e.target.value.length > 0)

        props.onChange && props.onChange(e)
    }

    const classes = ["text-input"]
    if (filled) {
        classes.push("has-content")
    }
    if (props.light === true) {
        classes.push("text-input-light")
    }

    const inputProps = {...props}
    delete inputProps.light
    delete inputProps.containerProps

    return (
        <div className="input-effect" ref={idRef} {...(props.containerProps || {})}>
            <input {...inputProps} id={id} onChange={onChange} className={classes.join(" ")} type="text" placeholder="" />
            <label htmlFor={id}>{props.placeholder}</label>
            <span className="focus-border"><i></i></span>
        </div>
    )
}