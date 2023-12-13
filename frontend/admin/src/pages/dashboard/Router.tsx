
import { Route, Routes } from "react-router-dom";
import UsersPage from "../users/UsersPage";

function Router() {
    return (
        <Routes>
            <Route path="/users" element={<UsersPage />} />
            <Route path="/" element={<p>HOME</p>} />
        </Routes>
    )
}

export default Router