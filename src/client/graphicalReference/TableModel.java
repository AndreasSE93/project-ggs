package graphicalReference;

import javax.swing.table.AbstractTableModel;

public class TableModel extends AbstractTableModel {
	/**
	 * 
	 */

	private static final long serialVersionUID = 1L;
	private String[] columnNames;
	private Object[][] data;

	public TableModel(Object[][] data, String[] columnNames) {
		this.columnNames = columnNames;
		this.data = data;
	}

	@Override
	public int getColumnCount() {
		return columnNames.length;
	}

	@Override
	public String getColumnName(int col) {
		return columnNames[col];
	}

	@Override
	public int getRowCount() {
		return data.length;

	}
	public void addAll(String[][] L) {
		data = new String[L.length][4];

		for (int i = 0; i < data.length; i++) {
			data[i][0] = L[i][0];
			data[i][1] = L[i][1];
			data[i][2] = L[i][2];
			data[i][3] = L[i][3];
		}
	}

	public void removeAll() {
		for (int i = 0; i < data.length; i++) {
			data[i][0] = null;
			data[i][1] = null;
			data[i][2] = null;
			data[i][3] = null;
		}
	}

	@Override
	public Object getValueAt(int rowIndex, int columnIndex) {
		return data[rowIndex][columnIndex];
	}

	@Override
	public boolean isCellEditable(int row, int col) {
		return false;
	}

}
