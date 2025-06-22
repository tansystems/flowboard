import * as React from 'react';
import { Edit, SimpleForm, TextInput } from 'react-admin';

export const TagEdit = props => (
    <Edit {...props} title="Редактировать тег">
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="name" label="Название" />
        </SimpleForm>
    </Edit>
); 