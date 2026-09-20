import { type SVGProps } from 'react'
import { ArrowIcon as ArrowIconComponent } from './arrow/arrow'
import { BellOutlineIcon as BellOutlineIconComponent } from './bell/bell-outline'
import { BellFilledIcon as BellFilledIconComponent } from './bell/bell-filled'
import { BoardOutlineIcon as BoardOutlineIconComponent } from './board/board-outline'
import { BoardFilledIcon as BoardFilledIconComponent } from './board/board-filled'
import { BugOutlineIcon as BugOutlineIconComponent } from './bug/bug-outline'
import { BugFilledIcon as BugFilledIconComponent } from './bug/bug-filled'
import { CheckboxIcon as CheckboxIconComponent } from './checkbox/checkbox'
import { CompletedOutlineIcon as CompletedOutlineIconComponent } from './completed/completed-outline'
import { CompletedFilledIcon as CompletedFilledIconComponent } from './completed/completed-filled'
import { CompletedTaskOutlineIcon as CompletedTaskOutlineIconComponent } from './completedTask/completedTask-outline'
import { CompletedTaskFilledIcon as CompletedTaskFilledIconComponent } from './completedTask/completedTask-filled'
import { CrossIcon as CrossIconComponent } from './cross/cross'
import { DocumentOutlineIcon as DocumentOutlineIconComponent } from './document/document-outline'
import { DocumentFilledIcon as DocumentFilledIconComponent } from './document/document-filled'
import { EyeIcon as EyeIconComponent } from './eye/eye'
import { EyeOffIcon as EyeOffIconComponent } from './eyeOff/eyeOff'
import { FilterOutlineIcon as FilterOutlineIconComponent } from './filter/filter-outline'
import { FilterFilledIcon as FilterFilledIconComponent } from './filter/filter-filled'
import { FolderOutlineIcon as FolderOutlineIconComponent } from './folder/folder-outline'
import { FolderFilledIcon as FolderFilledIconComponent } from './folder/folder-filled'
import { FoldersOutlineIcon as FoldersOutlineIconComponent } from './folders/folders-outline'
import { FoldersFilledIcon as FoldersFilledIconComponent } from './folders/folders-filled'
import { GitHubIcon as GitHubIconComponent } from './github/github'
import { GoogleIcon as GoogleIconComponent } from './google/google'
import { HomeOutlineIcon as HomeOutlineIconComponent } from './home/home-outline'
import { HomeFilledIcon as HomeFilledIconComponent } from './home/home-filled'
import { KeyOutlineIcon as KeyOutlineIconComponent } from './key/key-outline'
import { KeyFilledIcon as KeyFilledIconComponent } from './key/key-filled'
import { LockIcon as LockIconComponent } from './lock/lock'
import { Lock2Icon as Lock2IconComponent } from './lock2/lock2'
import { LogoIcon as LogoIconComponent } from './logo/logo'
import { MailIcon as MailIconComponent } from './mail/mail'
import { MegaphoneOutlineIcon as MegaphoneOutlineIconComponent } from './megaphone/megaphone-outline'
import { MegaphoneFilledIcon as MegaphoneFilledIconComponent } from './megaphone/megaphone-filled'
import { MilestoneIcon as MilestoneIconComponent } from './milestone/milestone'
import { MoonOutlineIcon as MoonOutlineIconComponent } from './moon/moon-outline'
import { MoonFilledIcon as MoonFilledIconComponent } from './moon/moon-filled'
import { PersonOutlineIcon as PersonOutlineIconComponent } from './person/person-outline'
import { PersonFilledIcon as PersonFilledIconComponent } from './person/person-filled'
import { PlusIcon as PlusIconComponent } from './plus/plus'
import { SearchIcon as SearchIconComponent } from './search/search'
import { SettingsOutlineIcon as SettingsOutlineIconComponent } from './settings/settings-outline'
import { SettingsFilledIcon as SettingsFilledIconComponent } from './settings/settings-filled'
import { SortIcon as SortIconComponent } from './sort/sort'
import { SpinnerIcon as SpinnerIconComponent } from './spinner/spinner'
import { StatisticUpIcon as StatisticUpIconComponent } from './statistic/statistic'
import { TaskOutlineIcon as TaskOutlineIconComponent } from './task/task-outline'
import { TaskFilledIcon as TaskFilledIconComponent } from './task/task-filled'
import { TeamOutlineIcon as TeamOutlineIconComponent } from './team/team-outline'
import { TeamFilledIcon as TeamFilledIconComponent } from './team/team-filled'
import { BookIcon } from './book/book'
import { HeadphonesIcon } from './headphones/headphones'
import { Arrow2Icon } from './arrow2/arrow2'
import { UploadIcon } from './upload/upload'
import { PaletteIcon } from './pallete/palette'
import { SchemaIcon } from './schema/schema'
import { ListIcon } from './list/list'
import { WidgetsIcon } from './widgets/widgets'
import { OptionsIcon } from './options/options'
import { TrashIcon } from './trash/trash'
import { SmileIcon } from './smile/smile'
import { SendIcon } from './send/send'
import { MessageIcon } from './message/message'
import { ShieldIcon } from './shield/shield'
import { SunIcon } from './sun/sun'
import { RefreshIcon } from './refresh/refresh'
import { ComputerIcon } from './computer/computer'
import { AtSignIcon } from './atsign/atsign'
import { EditIcon } from './edit/edit'

type IconProps = SVGProps<SVGSVGElement>

class Icons {
    static Headphones(props: IconProps) { return <HeadphonesIcon {...props}/> }
    static Book(props: IconProps) { return <BookIcon {...props}/> }
    static Arrow(props: IconProps) { return <ArrowIconComponent {...props} /> }
    static BellOutline(props: IconProps) { return <BellOutlineIconComponent {...props} /> }
    static BellFilled(props: IconProps) { return <BellFilledIconComponent {...props} /> }
    static BoardOutline(props: IconProps) { return <BoardOutlineIconComponent {...props} /> }
    static BoardFilled(props: IconProps) { return <BoardFilledIconComponent {...props} /> }
    static BugOutline(props: IconProps) { return <BugOutlineIconComponent {...props} /> }
    static BugFilled(props: IconProps) { return <BugFilledIconComponent {...props} /> }
    static Checkbox(props: IconProps) { return <CheckboxIconComponent {...props} /> }
    static CompletedOutline(props: IconProps) { return <CompletedOutlineIconComponent {...props} /> }
    static CompletedFilled(props: IconProps) { return <CompletedFilledIconComponent {...props} /> }
    static CompletedTaskOutline(props: IconProps) { return <CompletedTaskOutlineIconComponent {...props} /> }
    static CompletedTaskFilled(props: IconProps) { return <CompletedTaskFilledIconComponent {...props} /> }
    static Cross(props: IconProps) { return <CrossIconComponent {...props} /> }
    static DocumentOutline(props: IconProps) { return <DocumentOutlineIconComponent {...props} /> }
    static DocumentFilled(props: IconProps) { return <DocumentFilledIconComponent {...props} /> }
    static Eye(props: IconProps) { return <EyeIconComponent {...props} /> }
    static EyeOff(props: IconProps) { return <EyeOffIconComponent {...props} /> }
    static FilterOutline(props: IconProps) { return <FilterOutlineIconComponent {...props} /> }
    static FilterFilled(props: IconProps) { return <FilterFilledIconComponent {...props} /> }
    static FolderOutline(props: IconProps) { return <FolderOutlineIconComponent {...props} /> }
    static FolderFilled(props: IconProps) { return <FolderFilledIconComponent {...props} /> }
    static FoldersOutline(props: IconProps) { return <FoldersOutlineIconComponent {...props} /> }
    static FoldersFilled(props: IconProps) { return <FoldersFilledIconComponent {...props} /> }
    static Schema(props: IconProps) { return <SchemaIcon {...props} /> }
    static GitHub(props: IconProps) { return <GitHubIconComponent {...props} /> }
    static Palette(props: IconProps) { return <PaletteIcon {...props} /> }
    static Google(props: IconProps) { return <GoogleIconComponent {...props} /> }
    static HomeOutline(props: IconProps) { return <HomeOutlineIconComponent {...props} /> }
    static HomeFilled(props: IconProps) { return <HomeFilledIconComponent {...props} /> }
    static KeyOutline(props: IconProps) { return <KeyOutlineIconComponent {...props} /> }
    static KeyFilled(props: IconProps) { return <KeyFilledIconComponent {...props} /> }
    static Lock(props: IconProps) { return <LockIconComponent {...props} /> }
    static Lock2(props: IconProps) { return <Lock2IconComponent {...props} /> }
    static Logo(props: IconProps) { return <LogoIconComponent {...props} /> }
    static Mail(props: IconProps) { return <MailIconComponent {...props} /> }
    static MegaphoneOutline(props: IconProps) { return <MegaphoneOutlineIconComponent {...props} /> }
    static MegaphoneFilled(props: IconProps) { return <MegaphoneFilledIconComponent {...props} /> }
    static Milestone(props: IconProps) { return <MilestoneIconComponent {...props} /> }
    static MoonOutline(props: IconProps) { return <MoonOutlineIconComponent {...props} /> }
    static MoonFilled(props: IconProps) { return <MoonFilledIconComponent {...props} /> }
    static PersonOutline(props: IconProps) { return <PersonOutlineIconComponent {...props} /> }
    static PersonFilled(props: IconProps) { return <PersonFilledIconComponent {...props} /> }
    static Plus(props: IconProps) { return <PlusIconComponent {...props} /> }
    static Search(props: IconProps) { return <SearchIconComponent {...props} /> }
    static SettingsOutline(props: IconProps) { return <SettingsOutlineIconComponent {...props} /> }
    static SettingsFilled(props: IconProps) { return <SettingsFilledIconComponent {...props} /> }
    static Sort(props: IconProps) { return <SortIconComponent {...props} /> }
    static Spinner(props: IconProps) { return <SpinnerIconComponent {...props} /> }
    static StatisticUp(props: IconProps) { return <StatisticUpIconComponent {...props} /> }
    static TaskOutline(props: IconProps) { return <TaskOutlineIconComponent {...props} /> }
    static TaskFilled(props: IconProps) { return <TaskFilledIconComponent {...props} /> }
    static TeamOutline(props: IconProps) { return <TeamOutlineIconComponent {...props} /> }
    static TeamFilled(props: IconProps) { return <TeamFilledIconComponent {...props} /> }
    static Arrow2(props: IconProps) { return <Arrow2Icon {...props} /> }
    static Upload(props: IconProps) { return <UploadIcon {...props} /> }
    static List(props: IconProps) { return <ListIcon {...props} /> }
    static Widgets(props: IconProps) { return <WidgetsIcon {...props} /> }
    static Options(props: IconProps) { return <OptionsIcon {...props} /> }
    static Trash(props: IconProps) { return <TrashIcon {...props} /> }
    static Smile(props: IconProps) { return <SmileIcon {...props} /> }
    static Send(props: IconProps) { return <SendIcon {...props} /> }
    static Message(props: IconProps) { return <MessageIcon {...props} />  }
    static Shield(props: IconProps) { return <ShieldIcon {...props} /> }
    static Sun(props: IconProps) { return <SunIcon {...props} /> }
    static Refresh(props: IconProps) { return <RefreshIcon {...props} /> }
    static Computer(props: IconProps) { return <ComputerIcon {...props} /> }
    static AtSign(props: IconProps) { return <AtSignIcon {...props} /> }
    static Edit(props: IconProps) { return <EditIcon {...props} /> }
}

export default Icons