import React, { useEffect, useState } from 'react';

function App() {
 /* const [data, setData] = useState(null);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8080');
      const jsonData = await response.json();
      setData(jsonData);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  if (data === null) {
    return <div>Loading...</div>;
  }
  */

  /*
      <ul>
        {data.map((item) => (
          <li>{item.hello}</li>
        ))}
      </ul>
    */


  // Render your component with the fetched data
  return (
    <div>
      <h1>{process.env.REACT_APP_HELLO}</h1>
    </div>
  );
}

export default App;
