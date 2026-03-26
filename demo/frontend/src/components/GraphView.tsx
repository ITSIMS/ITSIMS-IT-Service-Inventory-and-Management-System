import { useState, useEffect } from 'react';
import { DependencyGraph, GraphNode } from '../types/service';
import { getDependencyGraph } from '../api/services';

export function GraphView() {
  const [graph, setGraph] = useState<DependencyGraph | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const loadGraph = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getDependencyGraph();
      setGraph(data);
    } catch {
      setError('Не удалось загрузить граф зависимостей');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadGraph();
  }, []);

  if (loading) {
    return <p style={{ color: '#6c757d', textAlign: 'center' }}>Загрузка графа...</p>;
  }

  if (error) {
    return <div className="error-banner">{error}</div>;
  }

  if (!graph || graph.nodes.length === 0) {
    return (
      <div className="graph-view graph-view--empty">
        <p>Зависимости не найдены</p>
        <button className="btn-primary" onClick={loadGraph}>
          Обновить граф
        </button>
      </div>
    );
  }

  // Build node map
  const nodeMap: Record<string, GraphNode> = {};
  for (const node of graph.nodes) {
    nodeMap[node.id] = node;
  }

  // Build adjacency: service_id -> list of depends_on nodes
  const adj: Record<string, string[]> = {};
  for (const edge of graph.edges) {
    if (!adj[edge.service_id]) {
      adj[edge.service_id] = [];
    }
    adj[edge.service_id].push(edge.depends_on_id);
  }

  return (
    <div className="graph-view">
      <div className="graph-view__header">
        <h2>Граф зависимостей</h2>
        <button className="btn-primary" onClick={loadGraph}>
          Обновить граф
        </button>
      </div>

      <div className="graph-view__summary">
        <span>{graph.nodes.length} узлов</span>
        <span>{graph.edges.length} связей</span>
      </div>

      <table className="graph-table">
        <thead>
          <tr>
            <th>Сервис</th>
            <th>Категория</th>
            <th>Статус</th>
            <th>Зависит от</th>
          </tr>
        </thead>
        <tbody>
          {graph.nodes.map((node) => {
            const deps = (adj[node.id] || [])
              .map((id) => nodeMap[id]?.name || id)
              .join(', ');
            return (
              <tr key={node.id}>
                <td className="graph-table__name">{node.name}</td>
                <td>{node.category}</td>
                <td>
                  <span className={`status-badge status-badge--${node.status}`}>
                    {node.status}
                  </span>
                </td>
                <td className="graph-table__deps">{deps || '—'}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
