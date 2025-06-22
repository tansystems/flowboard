import * as React from 'react';
import { Edit, SimpleForm, TextInput, ReferenceInput, SelectInput } from 'react-admin';

export const CommentEdit = props => (
    <Edit {...props} title="Редактировать комментарий">
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <ReferenceInput source="deal_id" reference="deals" label="Сделка">
                <SelectInput optionText="title" />
            </ReferenceInput>
            <ReferenceInput source="user_id" reference="users" label="Пользователь">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <TextInput source="content" label="Комментарий" />
        </SimpleForm>
    </Edit>
); 