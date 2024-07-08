export async function getUsers() {
    try {
        const response = await fetch('https://127.0.0.1:8443/listUsers', { // edit url to match backend
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
        });
        if (!response.ok) {
            throw new Error('Failed to get user');
        }
        return await response.json();
    } catch (err) {
        console.error(err);
        throw err;
    }
}