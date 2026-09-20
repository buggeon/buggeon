interface RadioButtonProps {
    value : boolean
    onChange: () => void
}

function RadioButton({value, onChange} : RadioButtonProps) {
    return(
        <button onClick={onChange} style={{
            all: 'unset',
            boxSizing: "border-box",
            width: 20, 
            height: 20, 
            borderRadius: '100%',
            padding: 0,
            backgroundColor: value ? "white" : "transparent",
            borderStyle: "solid",
            borderWidth: value ? 5 : 0.5,
            borderColor: value ? "var(--color-accent)" : "var(--color-bg)",
            cursor: 'pointer',
            transition: 'none',
            transform: 'none',
        }}>

        </button>
    )
}

export default RadioButton