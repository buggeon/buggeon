
import MainLayout from '../../layouts/mainLayout/mainLayout';
import '@excalidraw/excalidraw/index.css'

function SchemasScreen() {

    //const { projectId } = useParams()

    return(
        <MainLayout
            title="Project schemas"
            description="Here'are all project's schemas."
        >
            <img style={{width: 0, height: 250}} src='https://miro.com/api/v1/boards/uXjVHoTCOzk=/picture?etag=R3458764518193803894_1_20250610&size=180'/>
        </MainLayout>
    )
}

export default SchemasScreen;