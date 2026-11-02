import { Navigate, Route, Routes } from 'react-router-dom'
import EntryScreen from './screens/entry/entry'
import { useEffect, useState } from 'react'
import MainLayout from './layouts/mainLayout/mainLayout'
import ProjectsScreen from './screens/projects/projects'
import ProfileScreen from './screens/profile/profile'
import NotFoundScreen from './screens/notFound/notFound'
import SchemasScreen from './screens/schemas/schemas'
import BoardsScreen from './screens/boards/boards'
import CardAboutScreen from './screens/cardAbout/cardAbout'
import SettingsScreen from './screens/accountSettings/accountSettings'
import AccountSettingsScreen from './screens/accountSettings/accountSettings'
import ProjectSettingsScreen from './screens/projectSettings/projectSettings'
import MembersScreen from './screens/members/members'
import MemberAboutScreen from './screens/members/memberAbout/memberAbout'
import DiagramsScreen from './screens/diagrams/diagrams'

function App() {

    const [isAuth, setAuth] = useState(false)

    useEffect(() => {
        const accessToken = localStorage.getItem("accessToken")

        if(accessToken) {
            setAuth(true)
        }
    }, [])

    return (
        <Routes>
            <Route path='/' element={<Navigate to={isAuth ? "/dashboard" : "/auth"} replace/>}/>
            <Route path='/dashboard'>
                <Route index element={<Navigate to="profile" replace/>}/>
                <Route path="profile" element={<ProfileScreen/>}/>
                <Route path="projects" element={<ProjectsScreen/>}/>
                <Route path="settings" element={<AccountSettingsScreen/>}/>
            </Route>
            <Route path="/auth" element={<EntryScreen/>}/>

            <Route path="/projects/:projectId">
                <Route index element={<Navigate to="overview" replace/>}/>
                <Route path="overview" element={<MainLayout title="" description=""><></></MainLayout>}/>
                <Route path="settings" element={<ProjectSettingsScreen/>}/>
                <Route path="members" element={<MembersScreen/>}/>
                <Route path="boards" element={<BoardsScreen/>}/>
                <Route path="members/:memberId" element={<MemberAboutScreen/>}/>
                <Route path='boards/:boardId/cards/:cardId' element={<CardAboutScreen/>}></Route>
                <Route path='diagrams' element={<DiagramsScreen/>}/>
            </Route>

            <Route path='*' element={<NotFoundScreen/>}/>
        </Routes>
    )
}

export default App
