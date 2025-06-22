export const isAdmin = () => {
    return localStorage.getItem('user_role') === 'admin';
};

export const setUserRole = (role) => {
    localStorage.setItem('user_role', role);
};

export const clearUserRole = () => {
    localStorage.removeItem('user_role');
}; 