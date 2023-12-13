import { useState } from "react"
import UserContext from "../../context/UserContext"

function UsersPage() {
    const list = UserContext(state => state.list)
    const fetchList = UserContext(state => state.fetchList)
    const saveUser = UserContext(state => state.create)
    const [values, setValues] = useState({
        name: "",
        email: "",
        password: ""
    })
    console.log(list)
    function onSubmitHandler(e) {
        e.preventDefault()
        saveUser(values)
    }
    function onCahngeHandler(e) {
        e.preventDefault()
        setValues(current => ({
            ...current,
            [e.target.name]: e.target.value
        }))

    }
    return (
        <div>
            <h1>UsersPage</h1>
            <button onClick={fetchList}>Get List</button>
            <form onSubmit={onSubmitHandler}>
                <input type="text" name="name" placeholder="name" value={values.name} onChange={onCahngeHandler} />
                <input type="text" name="email" placeholder="email" value={values.email} onChange={onCahngeHandler} />
                <input type="password" name="password" placeholder="password" value={values.password} onChange={onCahngeHandler} />
                <button type="submit">SAVE</button>
            </form>
        </div>
    )
}

export default UsersPage