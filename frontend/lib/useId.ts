import { Ref, useState } from "react"

export function useId(): [string, Ref<any>] {
    const [id, setId] = useState<string | null>(null)
    if (id != null) {
        return [id, (r) => {}]
    }

    function generatePath(e: HTMLElement, currentPath: string = ""): string {
        const parent = e.parentElement
        if (parent != null) {
            const index = Array.prototype.indexOf.call(parent.children, e)
            const newPath = `${parent.nodeName}-${index}_${currentPath}`
            return generatePath(parent, newPath)
        }
        return currentPath
    }

    return ["", (r: HTMLElement) => {
        const hashCode = (s: string): string => s.split('').reduce((a,b)=>{a=((a<<5)-a)+b.charCodeAt(0);return a&a},0).toString()
        r != null && setId(`uid${hashCode(generatePath(r))}`)
    }]
}