package graphicalReference;

import javax.swing.table.AbstractTableModel;

public class TableModel extends AbstractTableModel {
	/**
	 * 
	 */
	
	
		
	
	private static final long serialVersionUID = 1L;
	private String[] columnNames;
	private Object[][] data;

	public TableModel( Object[][] data, String[] columnNames){
		this.columnNames = columnNames;
		this.data= data;
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
	
	

	@Override
	public Object getValueAt(int rowIndex, int columnIndex) {
		return data[rowIndex][columnIndex];
	}
	@Override
	public boolean isCellEditable(int row, int col) {
        return false;
    }

}
