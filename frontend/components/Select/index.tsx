import React, { ChangeEvent, useState } from "react"
import { useId } from "../../lib/useId"

export default function Select(props: React.DetailedHTMLProps<React.SelectHTMLAttributes<HTMLSelectElement>, HTMLSelectElement>) {
    const [id, idRef] = useId()
    
    return (
        <div className="select-effect" ref={idRef}>
            <select {...props} id={id} className="select">
                {props.children}
            </select>
            <label htmlFor={id}>{props.placeholder}</label>
        </div>
    )
}