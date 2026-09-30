import { CheckboxIcon } from "../../assets/checkbox/checkbox";
import Icons from "../../assets/icons";
import styles from './checkbox.module.scss'

function Checkbox({isSelected, setSelected} : {isSelected : boolean, setSelected : any}) {

    return(
        <div onClick={setSelected} className={`${styles.main} ${isSelected ? styles.selected : ""}`}>
            {
                isSelected && <Icons.Check
                    width={20}
                    height={20}
                    color="white"
                />
            }
        </div>
    )

}

export default Checkbox