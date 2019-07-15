import * as React from "react";
import Home from "@material-ui/icons/Home";
import Subject from "@material-ui/icons/Subject";
import People from "@material-ui/icons/People";
import InsertDriveFile from "@material-ui/icons/InsertDriveFile";
import Settings from "@material-ui/icons/Settings";
import { Index } from "../containers/page/Index";
import { LoginPage } from "../containers/page/auth/Login";
import { LogoutPage } from "../containers/page/auth/Logout";
import {User} from "../containers/page/User";
import {Entry} from "../containers/page/Entry";
import {UserEdit} from "../containers/page/UserEdit";

export interface IRouteInfo {
    name: string;
    path?: string;
    component?: any;
    showLeftNav: boolean;
    permission: number[];
    icon?: any;
    children?: IRouteInfo[];
    link?: boolean;
}

export const ROUTES: IRouteInfo[] = [
    // 全般
    {
        name: "ホーム",
        path: "/",
        icon: <Home/>,
        component: Index,
        showLeftNav: true,
        permission: [],
    },
    {
        name: "ログイン",
        path: "/auth/login",
        component: LoginPage,
        showLeftNav: false,
        permission: [],
    },
    {
        name: "ログアウト",
        path: "/auth/logout",
        component: LogoutPage,
        showLeftNav: false,
        permission: [],
    },
    {
        name: "エントリー",
        path: "/entries",
        icon: <Subject/>,
        component: Entry,
        showLeftNav: true,
        permission: [],
    },
    {
        name: "ページ",
        path: "/pages",
        icon: <InsertDriveFile/>,
        component: Entry,
        showLeftNav: true,
        permission: [],
    },
    {
        name: "ユーザー",
        path: "/users",
        icon: <People/>,
        component: User,
        showLeftNav: true,
        permission: [1],
        link: true,
        children: [
            {
                name: "ユーザー編集",
                path: "/users/:id/edit",
                icon: <People/>,
                component: UserEdit,
                showLeftNav: false,
                permission: [1],
            },
        ],
    },
    {
        name: "設定",
        path: "/settings",
        icon: <Settings/>,
        showLeftNav: true,
        permission: [1],
        children: [
            {
                name: "サイト情報",
                path: "/settings/site",
                component: User,
                showLeftNav: true,
                permission: [1],
            },
            {
                name: "Google Analytics",
                path: "/settings/google-analytics",
                component: User,
                showLeftNav: true,
                permission: [1],
            },
        ]
    },
    // {
    //     name: "外部リンク (Google)",
    //     path: "https://google.com",
    //     icon: <Search/>,
    //     showLeftNav: true,
    //     permission: [],
    // },
];