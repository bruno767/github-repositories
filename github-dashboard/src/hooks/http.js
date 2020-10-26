import { useEffect, useState } from 'react';

export const useHttp = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);
    const [apiUrl, setApiUrl] = useState("");

    useEffect(() => {
        if (apiUrl.length > 0 ) {
            setLoading(true);

            fetch(apiUrl)
                .then(data => data.json())
                .then(data => {
                    if (data.error) {
                        throw(data.error);
                    }
                    setLoading(false);
                    setData(data);
                    return data;
                })
                .catch(error => {
                    setLoading(false);
                    console.error(error);
                });
        }
    }, [apiUrl]);

    return {loading, data, callApi: setApiUrl };
}
