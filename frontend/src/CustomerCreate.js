import * as React from 'react';
import { Create, SimpleForm, TextInput } from 'react-admin';

export const CustomerCreate = (props) => (
    <Create {...props} title="Создать клиента">
        <SimpleForm>
            <TextInput source="name" label="Имя" />
            <TextInput source="email" label="Email" />
            <TextInput source="phone" label="Телефон" />
            <TextInput source="company" label="Компания" />
        </SimpleForm>
    </Create>
); 