import * as React from 'react';
import { Admin, Resource, ListGuesser, ExportButton, TopToolbar } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import { fetchUtils } from 'react-admin';
import { CustomerList } from './CustomerList';
import { CustomerEdit } from './CustomerEdit';
import { CustomerCreate } from './CustomerCreate';
import { DealList } from './DealList';
import { DealEdit } from './DealEdit';
import { DealCreate } from './DealCreate';
import { StatusList } from './StatusList';
import { StatusEdit } from './StatusEdit';
import { StatusCreate } from './StatusCreate';
import { TagList } from './TagList';
import { TagEdit } from './TagEdit';
import { TagCreate } from './TagCreate';
import { UserList } from './UserList';
import { UserEdit } from './UserEdit';
import { UserCreate } from './UserCreate';
import { CommentList } from './CommentList';
import { CommentEdit } from './CommentEdit';
import { CommentCreate } from './CommentCreate';
import { isAdmin } from './helpers';
import russianMessages from 'ra-language-russian';
import { createTheme } from '@mui/material/styles';
import { Dashboard } from './Dashboard';
import { setUserRole, clearUserRole } from './helpers';

const apiUrl = 'http://localhost:8080';

// JWT авторизация
const authProvider = {
    login: async ({ email, password }) => {
        const request = new Request(`${apiUrl}/auth/login`, {
            method: 'POST',
            body: JSON.stringify({ email, password }),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        });
        const response = await fetch(request);
        if (!response.ok) {
            throw new Error('Ошибка авторизации');
        }
        const { token } = await response.json();
        localStorage.setItem('jwt', token);
        // Получаем роль пользователя
        const meReq = new Request(`${apiUrl}/auth/me`, {
            method: 'GET',
            headers: new Headers({ 'Authorization': `Bearer ${token}` }),
        });
        const meResp = await fetch(meReq);
        if (meResp.ok) {
            const user = await meResp.json();
            setUserRole(user.role);
        }
    },
    logout: () => {
        localStorage.removeItem('jwt');
        clearUserRole();
        return Promise.resolve();
    },
    checkAuth: () => {
        return localStorage.getItem('jwt') ? Promise.resolve() : Promise.reject();
    },
    getPermissions: () => Promise.resolve(),
    checkError: (error) => {
        if (error.status === 401 || error.status === 403) {
            localStorage.removeItem('jwt');
            clearUserRole();
            return Promise.reject();
        }
        return Promise.resolve();
    },
};

// Добавляем JWT в каждый запрос
const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    const token = localStorage.getItem('jwt');
    if (token) {
        options.headers.set('Authorization', `Bearer ${token}`);
    }
    return fetchUtils.fetchJson(url, options);
};

const dataProvider = simpleRestProvider(apiUrl, httpClient);

const i18nProvider = () => russianMessages;

const myTheme = createTheme({
    palette: {
        primary: { main: '#1976d2' },
        secondary: { main: '#388e3c' },
        background: { default: '#f4f6f8' },
    },
    typography: {
        fontFamily: 'Roboto, Arial, sans-serif',
    },
});

const ListActions = (props) => (
    <TopToolbar>
        <ExportButton {...props} />
    </TopToolbar>
);

export default function App() {
    return (
        <Admin
            dataProvider={dataProvider}
            authProvider={authProvider}
            i18nProvider={i18nProvider}
            theme={myTheme}
            dashboard={Dashboard}
        >
            <Resource name="customers" list={props => <CustomerList {...props} actions={<ListActions />} />} edit={CustomerEdit} create={CustomerCreate} />
            <Resource name="deals" list={props => <DealList {...props} actions={<ListActions />} />} edit={DealEdit} create={DealCreate} />
            <Resource name="statuses" list={props => <StatusList {...props} actions={<ListActions />} />} edit={StatusEdit} create={StatusCreate} />
            <Resource name="tags" list={props => <TagList {...props} actions={<ListActions />} />} edit={TagEdit} create={TagCreate} />
            <Resource name="users" list={props => <UserList {...props} actions={<ListActions />} />} edit={UserEdit} create={UserCreate} />
            <Resource name="comments" list={props => <CommentList {...props} actions={<ListActions />} />} edit={CommentEdit} create={CommentCreate} />
        </Admin>
    );
} 