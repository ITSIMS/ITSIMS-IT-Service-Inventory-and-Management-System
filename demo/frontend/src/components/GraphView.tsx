import { useState, useEffect, useRef, useCallback } from 'react';
import { DependencyGraph } from '../types/service';
import { getDependencyGraph } from '../api/services';

// ─── Цвета категорий ───────────────────────────────────────────────────────────
const CATEGORY_COLORS: Record<string, string> = {
  'CI/CD':               '#10b981',
  'Platform':            '#3b82f6',
  'Security':            '#ef4444',
  'Monitoring':          '#f59e0b',
  'Infrastructure':      '#8b5cf6',
  'Project Management':  '#06b6d4',
  'Documentation':       '#f97316',
};
const DEFAULT_COLOR = '#64748b';

function categoryColor(cat: string): string {
  return CATEGORY_COLORS[cat] ?? DEFAULT_COLOR;
}

// ─── Физика force-directed ─────────────────────────────────────────────────────
interface PhysNode {
  id: string;
  name: string;
  category: string;
  status: string;
  x: number;
  y: number;
  vx: number;
  vy: number;
}

interface PhysEdge {
  id: string;
  source: string;
  target: string;
}

function initNodes(nodes: DependencyGraph['nodes'], W: number, H: number): PhysNode[] {
  const cx = W / 2, cy = H / 2;
  const r = Math.min(W, H) * 0.35;
  return nodes.map((n, i) => {
    const angle = (2 * Math.PI * i) / nodes.length;
    return {
      ...n,
      x: cx + r * Math.cos(angle),
      y: cy + r * Math.sin(angle),
      vx: 0,
      vy: 0,
    };
  });
}

const REPULSION  = 8000;
const SPRING_K   = 0.04;
const REST_LEN   = 160;
const GRAVITY    = 0.003;
const DAMPING    = 0.82;
const NODE_R     = 28;

function tick(nodes: PhysNode[], edges: PhysEdge[], W: number, H: number): PhysNode[] {
  const cx = W / 2, cy = H / 2;
  const next = nodes.map(n => ({ ...n }));
  const idx = Object.fromEntries(next.map((n, i) => [n.id, i]));

  // Repulsion
  for (let i = 0; i < next.length; i++) {
    for (let j = i + 1; j < next.length; j++) {
      const dx = next[i].x - next[j].x;
      const dy = next[i].y - next[j].y;
      const d2 = dx * dx + dy * dy + 1;
      const f  = REPULSION / d2;
      const nx = (dx / Math.sqrt(d2)) * f;
      const ny = (dy / Math.sqrt(d2)) * f;
      next[i].vx += nx; next[i].vy += ny;
      next[j].vx -= nx; next[j].vy -= ny;
    }
  }

  // Spring attraction along edges
  for (const e of edges) {
    const si = idx[e.source], ti = idx[e.target];
    if (si === undefined || ti === undefined) continue;
    const dx = next[ti].x - next[si].x;
    const dy = next[ti].y - next[si].y;
    const d  = Math.sqrt(dx * dx + dy * dy) + 0.01;
    const f  = SPRING_K * (d - REST_LEN);
    const fx = (dx / d) * f;
    const fy = (dy / d) * f;
    next[si].vx += fx; next[si].vy += fy;
    next[ti].vx -= fx; next[ti].vy -= fy;
  }

  // Center gravity + integrate + damping + bounds
  for (const n of next) {
    n.vx += (cx - n.x) * GRAVITY;
    n.vy += (cy - n.y) * GRAVITY;
    n.vx *= DAMPING;
    n.vy *= DAMPING;
    n.x  += n.vx;
    n.y  += n.vy;
    n.x = Math.max(NODE_R + 4, Math.min(W - NODE_R - 4, n.x));
    n.y = Math.max(NODE_R + 4, Math.min(H - NODE_R - 4, n.y));
  }

  return next;
}

// ─── Компонент ────────────────────────────────────────────────────────────────
const SVG_W = 900;
const SVG_H = 600;
const TICK_STEPS = 120; // начальных шагов до первого рендера

export function GraphView() {
  const [rawGraph, setRawGraph]   = useState<DependencyGraph | null>(null);
  const [loading, setLoading]     = useState(true);
  const [error, setError]         = useState<string | null>(null);
  const [nodes, setNodes]         = useState<PhysNode[]>([]);
  const [edges, setEdges]         = useState<PhysEdge[]>([]);
  const [selected, setSelected]   = useState<string | null>(null);
  const [settled, setSettled]     = useState(false);
  const rafRef = useRef<number>(0);
  const nodesRef = useRef<PhysNode[]>([]);

  // ── Загрузка данных ──────────────────────────────────────────────────────────
  const loadGraph = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);
      setSettled(false);
      setSelected(null);
      const data = await getDependencyGraph();
      setRawGraph(data);
    } catch {
      setError('Не удалось загрузить граф зависимостей');
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => { loadGraph(); }, [loadGraph]);

  // ── Инициализация + прогрев симуляции ────────────────────────────────────────
  useEffect(() => {
    if (!rawGraph) return;

    const physEdges: PhysEdge[] = rawGraph.edges.map(e => ({
      id: e.id, source: e.service_id, target: e.depends_on_id,
    }));

    let ns = initNodes(rawGraph.nodes, SVG_W, SVG_H);
    // Прогреваем 120 шагов перед отрисовкой
    for (let i = 0; i < TICK_STEPS; i++) {
      ns = tick(ns, physEdges, SVG_W, SVG_H);
    }

    nodesRef.current = ns;
    setNodes([...ns]);
    setEdges(physEdges);
    setSettled(true);

    // Продолжаем симуляцию анимированно ещё ~180 шагов
    let step = 0;
    const animate = () => {
      if (step++ > 180) return;
      nodesRef.current = tick(nodesRef.current, physEdges, SVG_W, SVG_H);
      setNodes([...nodesRef.current]);
      rafRef.current = requestAnimationFrame(animate);
    };
    rafRef.current = requestAnimationFrame(animate);
    return () => cancelAnimationFrame(rafRef.current);
  }, [rawGraph]);

  // ── Drag ─────────────────────────────────────────────────────────────────────
  const dragging = useRef<{ id: string; ox: number; oy: number } | null>(null);

  const onNodePointerDown = (e: React.PointerEvent<SVGGElement>, id: string) => {
    e.currentTarget.setPointerCapture(e.pointerId);
    const svg = e.currentTarget.closest('svg')!.getBoundingClientRect();
    const scaleX = SVG_W / svg.width;
    const scaleY = SVG_H / svg.height;
    const node = nodesRef.current.find(n => n.id === id)!;
    dragging.current = {
      id,
      ox: (e.clientX - svg.left) * scaleX - node.x,
      oy: (e.clientY - svg.top)  * scaleY - node.y,
    };
  };

  const onSVGPointerMove = (e: React.PointerEvent<SVGSVGElement>) => {
    if (!dragging.current) return;
    const svg = e.currentTarget.getBoundingClientRect();
    const scaleX = SVG_W / svg.width;
    const scaleY = SVG_H / svg.height;
    const nx = (e.clientX - svg.left) * scaleX - dragging.current.ox;
    const ny = (e.clientY - svg.top)  * scaleY - dragging.current.oy;
    nodesRef.current = nodesRef.current.map(n =>
      n.id === dragging.current!.id ? { ...n, x: nx, y: ny, vx: 0, vy: 0 } : n,
    );
    setNodes([...nodesRef.current]);
  };

  const onSVGPointerUp = () => { dragging.current = null; };

  // ── Вычисляем хайлайты ───────────────────────────────────────────────────────
  const highlightEdges = new Set<string>();
  const highlightNodes = new Set<string>();
  if (selected) {
    highlightNodes.add(selected);
    for (const e of edges) {
      if (e.source === selected || e.target === selected) {
        highlightEdges.add(e.id);
        highlightNodes.add(e.source);
        highlightNodes.add(e.target);
      }
    }
  }

  // ── Arrow helper ─────────────────────────────────────────────────────────────
  function arrowPoints(sx: number, sy: number, tx: number, ty: number) {
    const dx = tx - sx, dy = ty - sy;
    const d  = Math.sqrt(dx * dx + dy * dy);
    if (d < 1) return { x1: sx, y1: sy, x2: tx, y2: ty };
    const ux = dx / d, uy = dy / d;
    return {
      x1: sx + ux * (NODE_R + 2),
      y1: sy + uy * (NODE_R + 2),
      x2: tx - ux * (NODE_R + 8),
      y2: ty - uy * (NODE_R + 8),
    };
  }

  const nodeMap = Object.fromEntries(nodes.map(n => [n.id, n]));

  // ── Легенда ──────────────────────────────────────────────────────────────────
  const usedCategories = [...new Set(nodes.map(n => n.category))].sort();

  // ── Render ───────────────────────────────────────────────────────────────────
  if (loading) {
    return (
      <div className="graph-wrapper graph-wrapper--loading">
        <div className="graph-spinner" />
        <p>Загрузка графа зависимостей…</p>
      </div>
    );
  }
  if (error) {
    return (
      <div className="graph-wrapper graph-wrapper--error">
        <p>{error}</p>
        <button className="btn-primary" onClick={loadGraph}>Повторить</button>
      </div>
    );
  }
  if (!rawGraph || rawGraph.nodes.length === 0) {
    return (
      <div className="graph-wrapper graph-wrapper--empty">
        <p>Зависимости не найдены</p>
        <button className="btn-primary" onClick={loadGraph}>Обновить</button>
      </div>
    );
  }

  return (
    <div className="graph-wrapper">
      {/* Заголовок */}
      <div className="graph-header">
        <div className="graph-header__left">
          <h2 className="graph-title">Граф зависимостей IT-сервисов</h2>
          <span className="graph-badge">{nodes.length} узлов</span>
          <span className="graph-badge">{edges.length} связей</span>
        </div>
        <button className="btn-primary" onClick={loadGraph}>↺ Обновить</button>
      </div>

      {/* Легенда */}
      <div className="graph-legend">
        {usedCategories.map(cat => (
          <span key={cat} className="graph-legend__item">
            <span
              className="graph-legend__dot"
              style={{ background: categoryColor(cat) }}
            />
            {cat}
          </span>
        ))}
        <span className="graph-legend__item">
          <span className="graph-legend__dot graph-legend__dot--inactive" />
          inactive
        </span>
      </div>

      {/* Подсказка */}
      {selected
        ? <p className="graph-hint">Кликните ещё раз или по пустому месту, чтобы снять выделение</p>
        : <p className="graph-hint">Кликните на узел, чтобы выделить его связи · Перетащите узел для перемещения</p>
      }

      {/* SVG граф */}
      <div className="graph-canvas-wrap">
        <svg
          className="graph-canvas"
          viewBox={`0 0 ${SVG_W} ${SVG_H}`}
          onPointerMove={onSVGPointerMove}
          onPointerUp={onSVGPointerUp}
          onPointerLeave={onSVGPointerUp}
          onClick={e => {
            if ((e.target as SVGElement).tagName === 'svg') setSelected(null);
          }}
          style={{ opacity: settled ? 1 : 0, transition: 'opacity .4s' }}
        >
          <defs>
            {/* Маркер стрелки — обычный */}
            <marker
              id="arrow"
              viewBox="0 0 10 10"
              refX="6" refY="5"
              markerWidth="6" markerHeight="6"
              orient="auto-start-reverse"
            >
              <path d="M 0 0 L 10 5 L 0 10 z" fill="#94a3b8" />
            </marker>
            {/* Маркер стрелки — выделенный */}
            <marker
              id="arrow-hl"
              viewBox="0 0 10 10"
              refX="6" refY="5"
              markerWidth="6" markerHeight="6"
              orient="auto-start-reverse"
            >
              <path d="M 0 0 L 10 5 L 0 10 z" fill="#3b82f6" />
            </marker>
            {/* Тень для узлов */}
            <filter id="shadow" x="-20%" y="-20%" width="140%" height="140%">
              <feDropShadow dx="0" dy="2" stdDeviation="3" floodOpacity="0.25" />
            </filter>
            {/* Тень выделенного узла */}
            <filter id="shadow-hl" x="-30%" y="-30%" width="160%" height="160%">
              <feDropShadow dx="0" dy="0" stdDeviation="8" floodColor="#3b82f6" floodOpacity="0.6" />
            </filter>
          </defs>

          {/* Рёбра */}
          {edges.map(e => {
            const s = nodeMap[e.source], t = nodeMap[e.target];
            if (!s || !t) return null;
            const { x1, y1, x2, y2 } = arrowPoints(s.x, s.y, t.x, t.y);
            const isHl = selected ? highlightEdges.has(e.id) : false;
            const dimmed = selected && !isHl;
            return (
              <line
                key={e.id}
                x1={x1} y1={y1} x2={x2} y2={y2}
                stroke={isHl ? '#3b82f6' : '#94a3b8'}
                strokeWidth={isHl ? 2.5 : 1.5}
                strokeOpacity={dimmed ? 0.15 : 1}
                markerEnd={isHl ? 'url(#arrow-hl)' : 'url(#arrow)'}
                style={{ transition: 'stroke-opacity .2s, stroke .2s' }}
              />
            );
          })}

          {/* Узлы */}
          {nodes.map(n => {
            const color  = categoryColor(n.category);
            const isInactive = n.status === 'inactive';
            const isSelected = n.id === selected;
            const isHl = selected ? highlightNodes.has(n.id) : false;
            const dimmed = selected && !isHl;

            return (
              <g
                key={n.id}
                transform={`translate(${n.x},${n.y})`}
                style={{ cursor: 'grab', opacity: dimmed ? 0.25 : 1, transition: 'opacity .2s' }}
                onClick={e => { e.stopPropagation(); setSelected(n.id === selected ? null : n.id); }}
                onPointerDown={(e: React.PointerEvent<SVGGElement>) => onNodePointerDown(e, n.id)}
              >
                {/* Внешнее кольцо для выделенного */}
                {isSelected && (
                  <circle
                    r={NODE_R + 7}
                    fill="none"
                    stroke="#3b82f6"
                    strokeWidth={2}
                    strokeDasharray="5 3"
                    style={{ animation: 'spin 4s linear infinite' }}
                  />
                )}

                {/* Основной круг */}
                <circle
                  r={NODE_R}
                  fill={isInactive ? '#e2e8f0' : color}
                  stroke={isSelected ? '#3b82f6' : isInactive ? '#94a3b8' : color}
                  strokeWidth={isSelected ? 3 : 2}
                  filter={isSelected ? 'url(#shadow-hl)' : 'url(#shadow)'}
                  style={{ transition: 'r .15s' }}
                />

                {/* Индикатор inactive */}
                {isInactive && (
                  <circle r={8} cx={NODE_R - 4} cy={-(NODE_R - 4)}
                    fill="#94a3b8" stroke="white" strokeWidth={1.5}
                  />
                )}

                {/* Категорийная полоска (дуга вверху) */}
                {!isInactive && (
                  <circle r={NODE_R} fill="none"
                    stroke="rgba(255,255,255,0.35)" strokeWidth={5}
                    strokeDasharray={`${NODE_R * 0.9} 999`}
                    strokeLinecap="round"
                  />
                )}

                {/* Текст — аббревиатура */}
                <text
                  textAnchor="middle"
                  dominantBaseline="central"
                  fontSize={isInactive ? 10 : 11}
                  fontWeight="700"
                  fill={isInactive ? '#64748b' : 'white'}
                  style={{ pointerEvents: 'none', userSelect: 'none' }}
                >
                  {abbreviate(n.name)}
                </text>

                {/* Подпись под кругом */}
                <text
                  y={NODE_R + 14}
                  textAnchor="middle"
                  fontSize={10}
                  fontWeight={isSelected ? '700' : '500'}
                  fill={isSelected ? '#1e40af' : '#334155'}
                  style={{ pointerEvents: 'none', userSelect: 'none' }}
                >
                  {n.name.length > 14 ? n.name.slice(0, 13) + '…' : n.name}
                </text>
              </g>
            );
          })}
        </svg>
      </div>

      {/* Панель деталей выбранного узла */}
      {selected && (() => {
        const n = nodeMap[selected];
        if (!n) return null;
        const outgoing = edges.filter(e => e.source === selected).map(e => nodeMap[e.target]?.name).filter(Boolean);
        const incoming = edges.filter(e => e.target === selected).map(e => nodeMap[e.source]?.name).filter(Boolean);
        return (
          <div className="graph-detail">
            <div className="graph-detail__header" style={{ borderLeft: `4px solid ${categoryColor(n.category)}` }}>
              <strong>{n.name}</strong>
              <span className={`status-badge status-badge--${n.status}`}>{n.status}</span>
              <span className="graph-detail__cat">{n.category}</span>
            </div>
            {outgoing.length > 0 && (
              <div className="graph-detail__row">
                <span className="graph-detail__label">→ Зависит от:</span>
                <span>{outgoing.join(', ')}</span>
              </div>
            )}
            {incoming.length > 0 && (
              <div className="graph-detail__row">
                <span className="graph-detail__label">← Используется в:</span>
                <span>{incoming.join(', ')}</span>
              </div>
            )}
            {outgoing.length === 0 && incoming.length === 0 && (
              <div className="graph-detail__row">
                <span className="graph-detail__label">Нет прямых зависимостей</span>
              </div>
            )}
          </div>
        );
      })()}
    </div>
  );
}

function abbreviate(name: string): string {
  const words = name.split(/[\s\-_]+/).filter(Boolean);
  if (words.length === 1) return name.slice(0, 3).toUpperCase();
  return words.slice(0, 2).map(w => w[0].toUpperCase()).join('');
}
