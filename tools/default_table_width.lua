function Table(el)
    local col_lengths = {}
    local total_length = 0
    local num_cols = #el.colspecs

    -- Initialize col_lengths with zeros
    for i = 1, num_cols do
        col_lengths[i] = 0
    end

    -- Calculate total length of content in each column
    if el.rows then
        for _, row in ipairs(el.rows) do
            if row.cells then
                for i, cell in ipairs(row.cells) do
                    if cell and i <= num_cols then
                        local cell_text = pandoc.utils.stringify(cell)
                        local cell_length = #cell_text
                        col_lengths[i] = col_lengths[i] + cell_length
                        total_length = total_length + cell_length
                    end
                end
            end
        end
    end

    -- Assign proportional widths based on content length
    for i, cs in ipairs(el.colspecs) do
        if total_length > 0 then
            local proportion = col_lengths[i] / total_length
            -- Ensure proportion is between 0 and 1
            cs[2] = math.max(0.05, math.min(proportion, 1.0))
        else
            -- Fallback to equal distribution if total is 0
            cs[2] = 1.0 / num_cols
        end
    end

    return el
end
