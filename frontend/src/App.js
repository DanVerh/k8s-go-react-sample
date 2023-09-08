import React, { useEffect, useState } from 'react';

function App() {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost//backend');
      const jsonData = await response.json();
      setData(jsonData);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  if (data === null) {
    return <div>Loading...</div>;
  // Render your component with the fetched data
  }

  return (
    <ul>
    {data.map((item) => (
      <li>{item.hello}</li>
    ))}
  </ul>
  );
}

export default App;
